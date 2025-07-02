import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../Styles/ActivityDetail.css";
import { useAuth } from '../hooks/useAuth';
import { useEnrollment } from '../hooks/useEnrollment.jsx';
import { enrollmentApi } from '../utils/enrollmentApi.jsx';
import { toast } from 'react-toastify';

const ActivityDetail = () => {
    const [activity, setActivity] = useState(null);
    const [error, setError] = useState(null);
    const [capacityInfo, setCapacityInfo] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    // Usar el hook personalizado para manejar inscripciones
    const { isEnrolled, isLoading, error: enrollmentError, toggleEnrollment } = useEnrollment(id);

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

    useEffect(() => {
        const fetchCapacityInfo = async () => {
            if (!isAuthenticated) return;

            try {
                const capacityData = await enrollmentApi.getActivityCapacity(id);
                setCapacityInfo(capacityData);
            } catch (err) {
                console.error('Error fetching capacity info:', err);
                // Si el error es de autenticación, redirigir al login
                if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                    navigate('/login');
                }
            }
        };

        fetchCapacityInfo();
    }, [id, isAuthenticated, navigate]);

    const handleEnrollmentToggle = async () => {
        if (!isAuthenticated) {
            navigate("/login");
            return;
        }

        try {
            await toggleEnrollment();
            const message = isEnrolled ? 'Inscripción cancelada exitosamente!' : 'Inscripción exitosa!';
            toast.success(message);

            // Actualizar la información de capacidad después de la acción
            const capacityData = await enrollmentApi.getActivityCapacity(id);
            setCapacityInfo(capacityData);
        } catch (err) {
            toast.error(err.message);
            setError(err.message);
            // Si el error es de autenticación, redirigir al login
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
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

    const getCapacityDisplay = () => {
        if (!capacityInfo) {
            return <p><strong>Capacidad:</strong> {activity.capacity} personas</p>;
        }

        const { available_spots } = capacityInfo;

        return (
            <div className="capacity-info">
                <p><strong>Capacidad:</strong> {activity.capacity} personas</p>
                <p><strong>Lugares disponibles:</strong> {available_spots}</p>
            </div>
        );
    };

    return (
        <div className="activity-detail-container">
            <button className="back-button" onClick={() => navigate('/activities')}>
                Volver a Actividades
            </button>

            {(error || enrollmentError) && (
                <div className="error-message">{error || enrollmentError}</div>
            )}

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
                        {getCapacityDisplay()}
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
                            onClick={handleEnrollmentToggle}
                            className={isEnrolled ? "cancel-button" : "enroll-button"}
                            disabled={isLoading || (capacityInfo && capacityInfo.available_spots === 0 && !isEnrolled)}
                            style={
                                capacityInfo && capacityInfo.available_spots === 0 && !isEnrolled
                                    ? { backgroundColor: '#ccc', color: '#888', cursor: 'not-allowed' }
                                    : {}
                            }
                        >
                            {isLoading
                                ? 'Procesando...'
                                : isEnrolled
                                    ? 'Cancelar Inscripción'
                                    : (capacityInfo && capacityInfo.available_spots === 0 ? 'Sin cupo disponible' : 'Inscribirse')}
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