import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "../Styles/Activities.css";

const Activities = () => {
    const navigate = useNavigate();
    const [activities, setActivities] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
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
            setActivities(data);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const handleEnrollment = async (activityId) => {
        if (!isloggedin) {
            navigate("/login");
            return;
        }

        try {
            const token = localStorage.getItem("token");
            const response = await fetch('http://localhost:8080/enrollments', {
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

    if (loading) {
        return <div className="loading">Cargando actividades...</div>;
    }

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

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
            </div>
            <div className="activities-grid">
                {activities.map((activity) => (
                    <div key={activity.id} className="activity-card">
                        <h3>{activity.name}</h3>
                        <p>{activity.description}</p>
                        <div className="activity-details">
                            <p><strong>Profesor:</strong> {activity.profesor}</p>
                            <p><strong>Categoría:</strong> {activity.category}</p>
                            <p><strong>Capacidad:</strong> {activity.capacity} personas</p>
                        </div>
                        <div className="schedule">
                            <h4>Horarios:</h4>
                            {activity.schedules.map((schedule, idx) => (
                                <p key={idx}>
                                    {schedule.day}: {schedule.start_time} - {schedule.end_time}
                                </p>
                            ))}
                        </div>
                        <button
                            onClick={() => handleEnrollment(activity.id)}
                            className={isloggedin ? "enroll-button" : "login-required-button"}
                        >
                            {isloggedin ? "Inscribirse" : "Iniciar sesión para inscribirse"}
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Activities;
