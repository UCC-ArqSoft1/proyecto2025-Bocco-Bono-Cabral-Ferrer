import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import "../Styles/Activities.css";

const Activities = () => {
    const [activities, setActivities] = useState([]);
    const [error, setError] = useState(null);
    const [searchTerm, setSearchTerm] = useState('');
    const [isSearching, setIsSearching] = useState(false);
    const navigate = useNavigate();
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
