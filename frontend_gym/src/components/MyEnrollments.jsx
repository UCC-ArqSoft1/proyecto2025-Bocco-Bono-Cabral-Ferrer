import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';
import { enrollmentApi } from '../utils/enrollmentApi.jsx';
import { toast } from 'react-toastify';
import '../Styles/MyEnrollments.css';

const MyEnrollments = () => {
    const [enrollments, setEnrollments] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [cancellingId, setCancellingId] = useState(null);
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    useEffect(() => {
        if (isAuthenticated && !isAdmin) {
            fetchEnrollments();
        }
    }, [isAuthenticated, isAdmin]);

    const fetchEnrollments = async () => {
        try {
            setLoading(true);
            const data = await enrollmentApi.getUserEnrollments();
            setEnrollments(Array.isArray(data) ? data : []);
        } catch (err) {
            setError(err.message);
            console.error('Error fetching enrollments:', err);
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
        } finally {
            setLoading(false);
        }
    };

    const handleCancelEnrollment = async (activityId) => {
        try {
            setCancellingId(activityId);
            await enrollmentApi.cancelEnrollment(activityId);
            toast.success('Inscripción cancelada exitosamente!');

            setEnrollments(prev => prev.filter(activity => activity.id !== activityId));
        } catch (err) {
            toast.error(err.message);
            console.error('Error cancelling enrollment:', err);
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
        } finally {
            setCancellingId(null);
        }
    };

    const handleViewActivity = (activityId) => {
        navigate(`/activity/${activityId}`);
    };

    if (!isAuthenticated) {
        return (
            <div className="my-enrollments-container">
                <div className="login-prompt">
                    <p>Debes iniciar sesión para ver tus inscripciones</p>
                    <button onClick={() => navigate('/login')} className="login-button">
                        Iniciar Sesión
                    </button>
                </div>
            </div>
        );
    }

    if (isAdmin) {
        return (
            <div className="my-enrollments-container">
                <div className="admin-message">
                    <p>Los administradores no pueden inscribirse en actividades</p>
                </div>
            </div>
        );
    }

    if (loading) {
        return (
            <div className="my-enrollments-container">
                <div className="loading">Cargando tus inscripciones...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="my-enrollments-container">
                <div className="error">Error: {error}</div>
            </div>
        );
    }

    return (
        <div className="my-enrollments-container">
            <div className="enrollments-header">
                <h2>Mis Inscripciones</h2>
            </div>

            {Array.isArray(enrollments) && enrollments.length === 0 ? (
                <div className="no-enrollments">
                    <p>No tienes inscripciones activas</p>
                </div>
            ) : (
                <div className="enrollments-grid">
                    {enrollments.map((activity) => (
                        <div key={activity.id} className="enrollment-card">
                            <div className="enrollment-content" onClick={() => handleViewActivity(activity.id)}>
                                <h3>{activity.name}</h3>
                                <div className="activity-details">
                                    <p><strong>Profesor:</strong> {activity.profesor}</p>
                                    <p><strong>Categoría:</strong> {activity.category}</p>
                                    <div className="schedule-info">
                                        <h4>Horarios:</h4>
                                        {Array.isArray(activity.schedules) && activity.schedules.map((schedule, index) => (
                                            <p key={index}>
                                                {schedule.day}: {schedule.start_time} - {schedule.end_time}
                                            </p>
                                        ))}
                                    </div>
                                </div>
                                <div className="view-details">
                                    Click para ver más detalles
                                </div>
                            </div>
                            <div className="enrollment-actions">
                                <button
                                    onClick={(e) => {
                                        e.stopPropagation();
                                        handleCancelEnrollment(activity.id);
                                    }}
                                    className="cancel-enrollment-button"
                                    disabled={cancellingId === activity.id}
                                >
                                    {cancellingId === activity.id ? 'Cancelando...' : 'Cancelar Inscripción'}
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}
export default MyEnrollments;