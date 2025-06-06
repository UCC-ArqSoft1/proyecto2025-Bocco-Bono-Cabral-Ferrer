import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../Styles/SearchActivities.css';

const SearchActivities = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [activities, setActivities] = useState([]);
    const [isSearching, setIsSearching] = useState(false);
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const isLoggedIn = localStorage.getItem("isLogin") === "true";

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
        } catch (err) {
            setError('Error al buscar actividades. Por favor, intente nuevamente.');
            console.error('Error:', err);
        } finally {
            setIsSearching(false);
        }
    };

    const handleEnrollment = async (activityId) => {
        try {
            const response = await fetch('http://localhost:8080/enrollments', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify({
                    activity_id: activityId,
                    user_id: localStorage.getItem('userId')
                })
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || 'Error al inscribirse en la actividad');
            }

            alert('¡Inscripción exitosa!');
            navigate('/activities');
        } catch (err) {
            alert(err.message);
        }
    };

    return (
        <section className="search-section">
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

                {error && <div className="search-error">{error}</div>}

                {activities.length > 0 && (
                    <div className="search-results">
                        {activities.map((activity) => (
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
                                    {isLoggedIn && (
                                        <button
                                            onClick={() => handleEnrollment(activity.id)}
                                            className="enroll-button"
                                        >
                                            Inscribirse
                                        </button>
                                    )}
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </section>
    );
};

export default SearchActivities; 