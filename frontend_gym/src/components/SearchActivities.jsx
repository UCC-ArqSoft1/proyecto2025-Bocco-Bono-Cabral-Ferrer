import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../Styles/SearchActivities.css';
import { useAuth } from '../hooks/useAuth';
import { enrollmentApi } from '../utils/enrollmentApi.jsx';

const SearchActivities = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [activities, setActivities] = useState([]);
    const [isSearching, setIsSearching] = useState(false);
    const [error, setError] = useState('');
    const [enrollmentStatus, setEnrollmentStatus] = useState({});
    const [loadingStates, setLoadingStates] = useState({});
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    const checkEnrollmentStatus = async (activityId) => {
        if (!isAuthenticated || isAdmin) return false;

        try {
            const isEnrolled = await enrollmentApi.checkEnrollment(activityId);
            return isEnrolled;
        } catch (err) {
            console.error('Error checking enrollment status:', err);
            // Si el error es de autenticación, redirigir al login
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
        }
        return false;
    };

    const handleSearch = async (e) => {
        e.preventDefault();
        if (!searchTerm.trim()) return;

        setIsSearching(true);
        setError('');

        try {
            const response = await fetch(`http://localhost:8080/activities/search?keyword=${encodeURIComponent(searchTerm)}`);
            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || 'Error al buscar actividades');
            }

            setActivities(data);

            // Verificar estado de inscripción para cada actividad
            if (isAuthenticated && !isAdmin) {
                const enrollmentPromises = data.map(async (activity) => {
                    const isEnrolled = await checkEnrollmentStatus(activity.id);
                    return { activityId: activity.id, isEnrolled };
                });

                const enrollmentResults = await Promise.all(enrollmentPromises);
                const enrollmentMap = {};
                enrollmentResults.forEach(result => {
                    enrollmentMap[result.activityId] = result.isEnrolled;
                });
                setEnrollmentStatus(enrollmentMap);
            }
        } catch (err) {
            setError('Error al buscar actividades. Por favor, intente nuevamente.');
            console.error('Error:', err);
        } finally {
            setIsSearching(false);
        }
    };

    const handleEnrollment = async (activityId) => {
        if (!isAuthenticated) {
            navigate('/login');
            return;
        }

        setLoadingStates(prev => ({ ...prev, [activityId]: true }));
        try {
            await enrollmentApi.enroll(activityId);
            alert('Inscripción exitosa!');
            setEnrollmentStatus(prev => ({ ...prev, [activityId]: true }));
        } catch (err) {
            setError(err.message);
            // Si el error es de autenticación, redirigir al login
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
        } finally {
            setLoadingStates(prev => ({ ...prev, [activityId]: false }));
        }
    };

    const handleCancelEnrollment = async (activityId) => {
        if (!isAuthenticated) {
            navigate('/login');
            return;
        }

        setLoadingStates(prev => ({ ...prev, [activityId]: true }));
        try {
            await enrollmentApi.cancelEnrollment(activityId);
            alert('Inscripción cancelada exitosamente!');
            setEnrollmentStatus(prev => ({ ...prev, [activityId]: false }));
        } catch (err) {
            setError(err.message);
            // Si el error es de autenticación, redirigir al login
            if (err.message.includes('sesión ha expirado') || err.message.includes('token de autenticación')) {
                navigate('/login');
            }
        } finally {
            setLoadingStates(prev => ({ ...prev, [activityId]: false }));
        }
    };

    return (
        <section className="search-section">
            <div className="search-container">
                <div className="search-header">
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
                    {isAdmin && (
                        <button
                            className="admin-button"
                            onClick={() => navigate('/admin/activities')}
                        >
                            Administrar Actividades
                        </button>
                    )}
                </div>

                {error && <div className="search-error">{error}</div>}

                {activities.length > 0 && (
                    <div className="search-results">
                        {activities.map((activity) => {
                            const isEnrolled = enrollmentStatus[activity.id] || false;
                            const isLoading = loadingStates[activity.id] || false;

                            return (
                                <div key={activity.id} className="activity-card">
                                    <div className="activity-image">
                                        <img
                                            src={activity.image_url || '/images/default-activity.jpg'}
                                            alt={activity.name}
                                            onError={(e) => {
                                                e.target.src = '/images/default-activity.jpg';
                                            }}
                                        />
                                    </div>
                                    <div className="activity-content">
                                        <h3>{activity.name}</h3>
                                        <p>{activity.description}</p>
                                        <div className="schedule">
                                            <h4>Horarios:</h4>
                                            {Array.isArray(activity.schedules) && activity.schedules.map((schedule, index) => (
                                                <p key={index}>
                                                    {schedule.day}: {schedule.start_time} - {schedule.end_time}
                                                </p>
                                            ))}
                                        </div>
                                        {isAuthenticated && (
                                            <button
                                                onClick={() => isEnrolled ? handleCancelEnrollment(activity.id) : handleEnrollment(activity.id)}
                                                className={isEnrolled ? "cancel-button" : "enroll-button"}
                                                disabled={isLoading}
                                            >
                                                {isLoading ? 'Procesando...' : (isEnrolled ? 'Cancelar Inscripción' : 'Inscribirse')}
                                            </button>
                                        )}
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                )}
            </div>
        </section>
    );
};

export default SearchActivities; 