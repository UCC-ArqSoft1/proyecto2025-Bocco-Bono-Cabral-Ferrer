import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../Styles/ActivityDetail.css";
import { useAuth } from '../hooks/useAuth';
import { toast } from 'react-toastify';

const ActivityDetail = () => {
    const [activity, setActivity] = useState(null);
    const [error, setError] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

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
        if (!isAuthenticated) {
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

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || 'Error al inscribirse en la actividad');
            }

            toast.success('Inscripción exitosa!');
        } catch (err) {
            toast.error(err.message);
            setError(err.message);
        }
    };

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

            {error && <div className="error-message">{error}</div>}

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

                    {isAuthenticated && !isAdmin && (
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