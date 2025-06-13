import "../Styles/home.css";
import { FaDumbbell, FaUserFriends, FaClock, FaHeart } from 'react-icons/fa';
import { useNavigate } from "react-router-dom";
import React from 'react';
import { useAuth } from '../hooks/useAuth';

const Home = () => {
    const navigate = useNavigate();
    const { isAuthenticated } = useAuth();

    const handleJoinNow = () => {
        navigate("/register");
    };

    return (
        <div className="home-page">
            {/* Sección Hero */}
            <section className="hero-section">
                <div className="hero-content">
                    <h1>Transforma Tu Vida</h1>
                    <p>Comienza tu viaje fitness hoy con el mejor equipo y entrenadores</p>
                    {!isAuthenticated && (
                        <button className="cta-button" onClick={handleJoinNow}>¡Únete Ahora!</button>
                    )}
                </div>
            </section>

            {/* Sección de Características */}
            <section className="features-section">
                <div className="content-wrapper">
                    <h2>¿Por Qué Elegirnos?</h2>
                    <div className="features-grid">
                        <div className="feature-card">
                            <FaDumbbell className="feature-icon" />
                            <h3>Equipo Moderno</h3>
                            <p>Instalaciones de última generación y equipamiento fitness actualizado</p>
                        </div>
                        <div className="feature-card">
                            <FaUserFriends className="feature-icon" />
                            <h3>Entrenadores Expertos</h3>
                            <p>Profesionales calificados para guiar tu entrenamiento</p>
                        </div>
                        <div className="feature-card">
                            <FaClock className="feature-icon" />
                            <h3>Horarios Flexibles</h3>
                            <p>Abierto 24/7 para adaptarnos a tu rutina</p>
                        </div>
                        <div className="feature-card">
                            <FaHeart className="feature-icon" />
                            <h3>Comunidad</h3>
                            <p>Únete a una comunidad fitness comprometida</p>
                        </div>
                    </div>
                </div>
            </section>

            {/* Sección de Membresías */}
            <section className="membership-section">
                <div className="content-wrapper">
                    <h2>Planes de Membresía</h2>
                    <div className="plans-container">
                        <div className="plan-card">
                            <h3>Básico</h3>
                            <p className="price">$29.99/mes</p>
                            <ul>
                                <li>Acceso a sala de musculación</li>
                                <li>Equipamiento básico</li>
                                <li>Acceso a vestuarios</li>
                            </ul>
                            <button className="plan-button" onClick={handleJoinNow}>Seleccionar Plan</button>
                        </div>
                        <div className="plan-card featured">
                            <h3>Premium</h3>
                            <p className="price">$49.99/mes</p>
                            <ul>
                                <li>Acceso total al gimnasio</li>
                                <li>Clases grupales incluidas</li>
                                <li>Sesión con entrenador personal</li>
                                <li>Acceso al spa</li>
                            </ul>
                            <button className="plan-button" onClick={handleJoinNow}>Seleccionar Plan</button>
                        </div>
                        <div className="plan-card">
                            <h3>Elite</h3>
                            <p className="price">$79.99/mes</p>
                            <ul>
                                <li>Acceso 24/7 al gimnasio</li>
                                <li>Todas las características premium</li>
                                <li>4 sesiones PT/mes</li>
                                <li>Consulta nutricional</li>
                            </ul>
                            <button className="plan-button" onClick={handleJoinNow}>Seleccionar Plan</button>
                        </div>
                    </div>
                </div>
            </section>

            {/* Sección de Testimonios */}
            <section className="testimonials-section">
                <div className="content-wrapper">
                    <h2>Lo Que Dicen Nuestros Miembros</h2>
                    <div className="testimonials-container">
                        <div className="testimonial-card">
                            <p>"¡El mejor gimnasio al que he ido! Entrenadores increíbles y gran comunidad."</p>
                            <h4>Sara Jiménez</h4>
                        </div>
                        <div className="testimonial-card">
                            <p>"Transformé mi vida en solo 6 meses. ¡Las instalaciones son excelentes!"</p>
                            <h4>Miguel Torres</h4>
                        </div>
                        <div className="testimonial-card">
                            <p>"Horarios flexibles y personal profesional. ¡No se puede pedir más!"</p>
                            <h4>Ana García</h4>
                        </div>
                    </div>
                </div>
            </section>

            {/* Sección de Llamada a la Acción */}
            <section className="cta-section">
                <div className="content-wrapper">
                    <h2>¿Listo para Comenzar tu Viaje Fitness?</h2>
                    <p>¡Únete ahora y obtén tu primer mes al 50% de descuento!</p>
                    {!isAuthenticated && (
                        <button className="cta-button" onClick={handleJoinNow}>¡Comienza Hoy!</button>
                    )}
                </div>
            </section>
        </div>
    );
}

export default Home; 