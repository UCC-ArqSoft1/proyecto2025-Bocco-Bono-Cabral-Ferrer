import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';
import MyEnrollments from '../components/MyEnrollments';
import "../Styles/Activities.css";

const Activities = () => {
    const [activities, setActivities] = useState([]);
    const [error, setError] = useState(null);
    const [searchTerm, setSearchTerm] = useState('');
    const [isSearching, setIsSearching] = useState(false);
    const [showMyEnrollments, setShowMyEnrollments] = useState(false);
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    useEffect(() => {
        if (!showMyEnrollments) {
            fetchActivities();
        }
    }, [showMyEnrollments]);

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

    const handleMyEnrollmentsClick = () => {
        setShowMyEnrollments(true);
    };

    const handleBackToActivities = () => {
        setShowMyEnrollments(false);
    };

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    if (showMyEnrollments) {
        return (
            <div className="activities-container">
                <div className="activities-header">
                    <button
                        className="back-to-activities-button"
                        onClick={handleBackToActivities}
                    >
                        ← Volver a Actividades
                    </button>
                    <h2></h2>
                </div>
                <MyEnrollments />
            </div>
        );
    }

    return (
        <div className="activities-container">
            <div className="activities-header">
                <h2>Actividades Disponibles</h2>
                <div className="header-buttons">
                    {isAuthenticated && !isAdmin && (
                        <button
                            className="my-enrollments-button"
                            onClick={handleMyEnrollmentsClick}
                        >
                            Mis Inscripciones
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

            <div className="activities-grid">
                {Array.isArray(activities) && activities.map((activity) => (
                    <div
                        key={activity.id}
                        className="activity-card"
                        onClick={() => navigate(`/activity/${activity.id}`)}
                    >
                        <h3>{activity.name}</h3>
                        <div className="activity-preview-details">
                            <p><strong>Profesor:</strong> {activity.profesor}</p>
                            <div className="schedule-preview">
                                <h4>Horarios:</h4>
                                {Array.isArray(activity.schedules) && activity.schedules.map((schedule, index) => (
                                    <p key={index}>
                                        {schedule.day}: {schedule.start_time} - {schedule.end_time}
                                    </p>
                                ))}
                            </div>
                        </div>
                        <div className="view-more">
                            Click para ver más detalles
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Activities;
