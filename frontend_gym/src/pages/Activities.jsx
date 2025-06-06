import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import "../Styles/Activities.css";

const Activities = () => {
    const [activities, setActivities] = useState([]);
    const [error, setError] = useState(null);
    const [searchTerm, setSearchTerm] = useState('');
    const [isSearching, setIsSearching] = useState(false);
    const navigate = useNavigate();
    const isloggedin = localStorage.getItem("isLogin") === "true";
    const userTypeId = parseInt(localStorage.getItem("userTypeId"), 10);
    const isAdmin = userTypeId === 1;

    useEffect(() => {
        fetchActivities();
    }, []);

    const fetchActivities = async () => {
        try {
            const response = await fetch('http://localhost:8080/activities');
            if (!response.ok) {
                throw new Error('Error al cargar las actividades');
            }
            const data = await response.json();
            console.log('Actividades cargadas:', data); // Para debug
            setActivities(data);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleSearch = async (e) => {
        e.preventDefault();
        if (!searchTerm.trim()) {
            fetchActivities();
            return;
        }

        setIsSearching(true);
        setError(null);

        try {
            const response = await fetch(`http://localhost:8080/activities/search?keyword=${encodeURIComponent(searchTerm)}`);
            if (!response.ok) {
                const data = await response.json();
                throw new Error(data.error || 'Error al buscar actividades');
            }
            const data = await response.json();
            setActivities(data);
        } catch (err) {
            setError('Error al buscar actividades. Por favor, intente nuevamente.');
            console.error('Error:', err);
        } finally {
            setIsSearching(false);
        }
    };

    const handleEnrollment = async (activityId) => {
        if (!isloggedin) {
            navigate("/login");
            return;
        }

        try {
            const token = localStorage.getItem("token");
            const response = await fetch('http://localhost:8080/enrollment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    activity_id: activityId
                })
            });

            const data = await response.json();

            if (response.ok) {
                alert('¡Inscripción exitosa!');
            } else {
                // Manejar errores específicos del backend
                if (data.error && typeof data.error === 'string') {
                    if (data.error.includes("already enrolled")) {
                        alert("Ya estás inscrito en esta actividad");
                    } else if (data.error.includes("full capacity")) {
                        alert("Lo sentimos, esta actividad está llena");
                    } else {
                        alert("Error al inscribirse: " + data.error);
                    }
                } else {
                    alert("Error desconocido al procesar la inscripción");
                }
            }
        } catch (error) {
            // Manejar errores de red o parsing JSON
            if (error instanceof TypeError && error.message === 'Failed to fetch') {
                alert("Error de conexión: No se pudo conectar con el servidor");
            } else if (error instanceof SyntaxError) {
                alert("Error del servidor: Respuesta inválida");
            } else {
                alert("Error inesperado: " + error.message);
            }
            console.error("Error detallado:", error);
        }
    };

    const handleDelete = async (activityId) => {
        if (!isAdmin) {
            alert('No tienes permisos para eliminar actividades');
            return;
        }

        if (!window.confirm('¿Estás seguro de que deseas eliminar esta actividad?')) {
            return;
        }

        try {
            const token = localStorage.getItem("token");
            const response = await fetch(`http://localhost:8080/activities/${activityId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.ok) {
                alert('Actividad eliminada con éxito');
                fetchActivities(); // Recargar la lista de actividades
            } else {
                const data = await response.json();
                alert(data.error || 'Error al eliminar la actividad');
            }
        } catch (error) {
            alert('Error al procesar la eliminación');
        }
    };

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    const getImageUrl = (imageUrl) => {
        if (!imageUrl) return null;
        if (imageUrl.startsWith('http')) return imageUrl;
        return imageUrl.startsWith('/') ? imageUrl : `/${imageUrl}`;
    };

    return (
        <div className="activities-container">
            <div className="activities-header">
                <h2>Actividades Disponibles</h2>
                {isAdmin && (
                    <button
                        className="admin-button"
                        onClick={() => navigate("/admin/activities")}
                    >
                        Administrar Actividades
                    </button>
                )}
                <div className="search-container">
                    <form onSubmit={handleSearch} className="search-form">
                        <input
                            type="text"
                            value={searchTerm}
                            onChange={(e) => setSearchTerm(e.target.value)}
                            placeholder="Buscar actividades..."
                            className="search-input"
                        />
                        <button type="submit" className="search-button" disabled={isSearching}>
                            {isSearching ? 'Buscando...' : 'Buscar'}
                        </button>
                    </form>
                </div>
            </div>
            <div className="activities-grid">
                {Array.isArray(activities) && activities.map((activity) => (
                    <div key={activity.id} className="activity-card">
                        {activity.image_url && (
                            <div className="activity-image">
                                <img
                                    src={getImageUrl(activity.image_url)}
                                    alt={activity.name}
                                    onError={(e) => {
                                        console.error('Error cargando imagen:', e.target.src);
                                        e.target.style.display = 'none';
                                    }}
                                />
                            </div>
                        )}
                        <h3>{activity.name}</h3>
                        <p>{activity.description}</p>
                        <div className="activity-details">
                            <p><strong>Profesor:</strong> {activity.profesor}</p>
                            <p><strong>Categoría:</strong> {activity.category}</p>
                            <p><strong>Capacidad:</strong> {activity.capacity} personas</p>
                        </div>
                        <div className="schedule">
                            <h4>Horarios:</h4>
                            {Array.isArray(activity.schedules) && activity.schedules.map((schedule, index) => (
                                <p key={index}>
                                    {schedule.day}: {schedule.start_time} - {schedule.end_time}
                                </p>
                            ))}
                        </div>
                        {isloggedin && !isAdmin && (
                            <button
                                onClick={() => handleEnrollment(activity.id)}
                                className={isloggedin ? "enroll-button" : "login-required-button"}
                            >
                                {isloggedin ? "Inscribirse" : "Iniciar sesión para inscribirse"}
                            </button>
                        )}
                        {isAdmin && (
                            <div className="admin-buttons">
                                <button
                                    onClick={() => navigate(`/edit-activity/${activity.id}`)}
                                >
                                    Editar
                                </button>
                                <button
                                    onClick={() => handleDelete(activity.id)}
                                    className="delete-button"
                                >
                                    Eliminar
                                </button>
                            </div>
                        )}
                    </div>
                ))}
            </div>
            {isAdmin && (
                <button
                    onClick={() => navigate('/admin/activities')}
                    className="create-button"
                >
                    Crear Nueva Actividad
                </button>
            )}
        </div>
    );
};

export default Activities;
