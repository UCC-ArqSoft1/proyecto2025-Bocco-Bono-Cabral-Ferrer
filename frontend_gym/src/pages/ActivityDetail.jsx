import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../Styles/ActivityDetail.css";

const ActivityDetail = () => {
    const [activity, setActivity] = useState(null);
    const [error, setError] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();
    const isloggedin = localStorage.getItem("isLogin") === "true";
    const userTypeId = parseInt(localStorage.getItem("userTypeId"), 10);
    const isAdmin = userTypeId === 1;

    useEffect(() => {
        const fetchActivity = async () => {
            try {
                const response = await fetch(`http://localhost:8080/activities/${id}`);
                if (!response.ok) {
                    throw new Error('Error al cargar la actividad');
                }
                const data = await response.json();
                setActivity(data);
            } catch (err) {
                setError(err.message);
            }
        };

        fetchActivity();
    }, [id]);

    const handleEnrollment = async () => {
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
                    activity_id: id
                })
            });

            let data;
            try {
                data = await response.json();
            } catch (e) {
                console.error('Error parsing response:', e);
                alert('Error al procesar la respuesta del servidor');
                return;
            }

            if (response.ok) {
                alert('¡Inscripción exitosa!');
            } else {
                if (data.error && typeof data.error === 'string') {
                    if (data.error.includes("sql") || data.error.includes("Scan error")) {
                        alert("Ya estás inscrito en esta actividad");
                    } else if (data.error.toLowerCase().includes("already enrolled") ||
                        data.error.toLowerCase().includes("ya inscrito") ||
                        data.error.toLowerCase().includes("ya está inscrito")) {
                        alert("Ya estás inscrito en esta actividad");
                    } else if (data.error.toLowerCase().includes("full capacity") ||
                        data.error.toLowerCase().includes("capacidad llena") ||
                        data.error.toLowerCase().includes("no hay cupo")) {
                        alert("Lo sentimos, esta actividad está llena");
                    } else {
                        alert("No se pudo completar la inscripción. Por favor, intenta nuevamente más tarde.");
                        console.error("Error original:", data.error);
                    }
                } else {
                    alert("Error desconocido al procesar la inscripción");
                }
            }
        } catch (error) {
            if (error instanceof TypeError && error.message === 'Failed to fetch') {
                alert("Error de conexión: No se pudo conectar con el servidor");
            } else {
                alert("Error al procesar la inscripción. Por favor, intenta nuevamente más tarde.");
            }
            console.error("Error detallado:", error);
        }
    };

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    if (!activity) {
        return <div className="loading">Cargando...</div>;
    }

    const getImageUrl = (imageUrl) => {
        if (!imageUrl) return null;
        if (imageUrl.startsWith('http')) return imageUrl;
        return imageUrl.startsWith('/') ? imageUrl : `/${imageUrl}`;
    };

    return (
        <div className="activity-detail-container">
            <button className="back-button" onClick={() => navigate('/activities')}>
                Volver a Actividades
            </button>

            <div className="activity-detail-card">
                {activity.image_url && (
                    <div className="activity-detail-image">
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

                <h1>{activity.name}</h1>

                <div className="activity-detail-content">
                    <div className="detail-section">
                        <h3>Descripción</h3>
                        <p>{activity.description}</p>
                    </div>

                    <div className="detail-section">
                        <h3>Detalles</h3>
                        <p><strong>Profesor:</strong> {activity.profesor}</p>
                        <p><strong>Categoría:</strong> {activity.category}</p>
                        <p><strong>Capacidad:</strong> {activity.capacity} personas</p>
                    </div>

                    <div className="detail-section">
                        <h3>Horarios</h3>
                        {Array.isArray(activity.schedules) && activity.schedules.map((schedule, index) => (
                            <p key={index}>
                                {schedule.day}: {schedule.start_time} - {schedule.end_time}
                            </p>
                        ))}
                    </div>

                    {isloggedin && !isAdmin && (
                        <button
                            onClick={handleEnrollment}
                            className="enroll-button"
                        >
                            Inscribirse
                        </button>
                    )}

                    {isAdmin && (
                        <button
                            className="admin-button"
                            onClick={() => navigate("/admin/activities")}
                        >
                            Administrar Actividades
                        </button>
                    )}
                </div>
            </div>
        </div>
    );
};

export default ActivityDetail; 