import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import "../Styles/activitiesAdmin.css";

const Admn = () => {
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    // Verificar si el usuario es administrador
    useEffect(() => {
        if (!isAuthenticated) {
            navigate('/login');
            return;
        }

        if (!isAdmin) {
            alert('No tienes permisos de administrador para acceder a esta página');
            navigate('/activities');
            return;
        }
    }, [isAuthenticated, isAdmin, navigate]);

    // Si no es admin, no renderizar nada
    if (!isAuthenticated || !isAdmin) {
        return null;
    }

    // ------------------------------
    // 1) ESTADOS
    // ------------------------------
    const [activities, setActivities] = useState([]);

    // Para controlar si estamos editando o creando
    const [isEditing, setIsEditing] = useState(false);
    const [editId, setEditId] = useState(null);

    // Campos del formulario básicos
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [capacity, setCapacity] = useState(0);
    const [category, setCategory] = useState("");
    const [profesor, setProfesor] = useState("");
    const [imageUrl, setImageUrl] = useState("");

    // Estado para horarios
    const [schedules, setSchedules] = useState([
        { day: "Lunes", start_time: "08:00", end_time: "09:00" }
    ]);
    const token = localStorage.getItem("token") || "";

    // ------------------------------
    // 2) FETCH DE ACTIVIDADES (GET /activities)
    // ------------------------------
    const fetchActivities = async () => {
        try {
            const resp = await fetch("http://localhost:8080/activities", {
                headers: {
                    "Content-Type": "application/json",
                    ...(token ? { Authorization: `Bearer ${token}` } : {}),
                },
            });
            if (!resp.ok) {
                const errorData = await resp.json();
                if (resp.status === 403) {
                    throw new Error("No tienes permisos de administrador para ver las actividades");
                }
                throw new Error(errorData.error || "Error al obtener actividades");
            }

            const data = await resp.json();
            setActivities(Array.isArray(data) ? data : []);
        } catch (error) {
            console.error("Error al cargar actividades:", error);
            alert(error.message || "No se pudieron cargar las actividades. Revisa la consola.");
        }
    };

    useEffect(() => {
        fetchActivities();
    }, []);

    // ------------------------------
    // 3) LIMPIAR FORMULARIO
    // ------------------------------
    const clearForm = () => {
        setIsEditing(false);
        setEditId(null);
        setName("");
        setDescription("");
        setCapacity(0);
        setCategory("");
        setProfesor("");
        setImageUrl("");
        setSchedules([{ day: "Lunes", start_time: "08:00", end_time: "09:00" }]);
    };

    // Funciones para manejar horarios
    const addSchedule = () => {
        setSchedules([...schedules, { day: "Lunes", start_time: "08:00", end_time: "09:00" }]);
    };

    const removeSchedule = (index) => {
        setSchedules(schedules.filter((_, i) => i !== index));
    };

    const updateSchedule = (index, field, value) => {
        const newSchedules = [...schedules];
        newSchedules[index] = { ...newSchedules[index], [field]: value };
        setSchedules(newSchedules);
    };

    // ------------------------------
    // 4) MANEJAR SUBMIT (CREAR / ACTUALIZAR)
    // ------------------------------
    const handleSubmit = async (e) => {
        e.preventDefault();

        // ----- Validaciones mínimas -----
        if (!name.trim() || !description.trim() || !category.trim() || !profesor.trim()) {
            alert("Los campos Nombre, Descripción, Categoría y Profesor son obligatorios.");
            return;
        }
        if (capacity <= 0) {
            alert("La capacidad debe ser un número entero positivo.");
            return;
        }

        if (schedules.length === 0) {
            alert("Debe agregar al menos un horario.");
            return;
        }

        // ----- Construir el body según el dominio -----
        const body = {
            name: name.trim(),
            description: description.trim(),
            capacity: parseInt(capacity, 10),
            category: category.trim(),
            profesor: profesor.trim(),
            image_url: imageUrl.trim(),
            schedules: schedules
        };

        try {
            let resp;
            if (isEditing && editId) {
                // --> UPDATE (PUT /activities/:id)
                resp = await fetch(`http://localhost:8080/activities/${editId}`, {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        ...(token ? { Authorization: `Bearer ${token}` } : {}),
                    },
                    body: JSON.stringify(body),
                });
                if (!resp.ok) {
                    const errorData = await resp.json();
                    if (resp.status === 403) {
                        throw new Error("No tienes permisos de administrador para actualizar actividades");
                    }
                    throw new Error(errorData.error || "Error al actualizar");
                }
                alert("Actividad actualizada correctamente.");
            } else {
                // --> CREATE (POST /activities)
                resp = await fetch("http://localhost:8080/activities", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        ...(token ? { Authorization: `Bearer ${token}` } : {}),
                    },
                    body: JSON.stringify(body),
                });
                if (!resp.ok) {
                    const errorData = await resp.json();
                    if (resp.status === 403) {
                        throw new Error("No tienes permisos de administrador para crear actividades");
                    }
                    throw new Error(errorData.error || "Error al crear");
                }
                alert("Actividad creada correctamente.");
            }

            clearForm();
            fetchActivities();
        } catch (err) {
            console.error(err);
            alert(err.message || "Ocurrió un problema con el servidor. Revisa la consola.");
        }
    };

    // ------------------------------
    // 5) ELIMINAR ACTIVIDAD (DELETE /activities/:id)
    // ------------------------------
    const handleDelete = async (id) => {
        if (!window.confirm("¿Estás seguro que deseas eliminar esta actividad?")) return;
        try {
            const resp = await fetch(`http://localhost:8080/activities/${id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    ...(token ? { Authorization: `Bearer ${token}` } : {}),
                },
            });
            if (!resp.ok) {
                const errorData = await resp.json();
                if (resp.status === 403) {
                    throw new Error("No tienes permisos de administrador para eliminar actividades");
                }
                throw new Error(errorData.error || "Error al eliminar");
            }
            alert("Actividad eliminada.");
            fetchActivities();
        } catch (err) {
            console.error(err);
            alert(err.message || "No se pudo eliminar la actividad. Revisa la consola.");
        }
    };

    // ------------------------------
    // 6) EDITAR: CARGAR DATOS AL FORMULARIO
    // ------------------------------
    const handleEdit = (activity) => {
        setIsEditing(true);
        setEditId(activity.id);
        setName(activity.name);
        setDescription(activity.description);
        setCapacity(activity.capacity);
        setCategory(activity.category);
        setProfesor(activity.profesor);
        setImageUrl(activity.image_url || "");
        setSchedules(activity.schedules);
    };

    const getImageUrl = (imageUrl) => {
        if (!imageUrl) return null;
        if (imageUrl.startsWith('http')) return imageUrl;
        return imageUrl.startsWith('/') ? imageUrl : `/${imageUrl}`;
    };

    // ------------------------------
    // 7) RENDER
    // ------------------------------
    return (
        <div className="admin-container">
            <h1>Panel de Administración de Actividades</h1>

            {/* ===========================================================
            FORMULARIO DE CREACIÓN / EDICIÓN (id="admin-form")
            =========================================================== */}
            <div id="admin-form" className="admin-form-card">
                <h2>{isEditing ? "Editar Actividad" : "Crear Nueva Actividad"}</h2>
                <form onSubmit={handleSubmit}>
                    {/* 1) Nombre */}
                    <label>
                        Nombre:
                        <input
                            type="text"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            placeholder="Ej. Yoga"
                        />
                    </label>

                    {/* 2) Descripción */}
                    <label>
                        Descripción:
                        <textarea
                            rows={3}
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                            placeholder="Ej. Clases de yoga para principiantes"
                        />
                    </label>

                    {/* 3) Capacidad */}
                    <label>
                        Capacidad:
                        <input
                            type="number"
                            min="1"
                            value={capacity}
                            onChange={(e) => setCapacity(e.target.value)}
                            placeholder="Ej. 20"
                        />
                    </label>

                    {/* 4) Categoría */}
                    <label>
                        Categoría:
                        <input
                            type="text"
                            value={category}
                            onChange={(e) => setCategory(e.target.value)}
                            placeholder="Ej. Deporte / Funcional / Pilates"
                        />
                    </label>

                    {/* 5) Profesor */}
                    <label>
                        Profesor:
                        <input
                            type="text"
                            value={profesor}
                            onChange={(e) => setProfesor(e.target.value)}
                            placeholder="Ej. Carlos Gómez"
                        />
                    </label>

                    {/* 6) URL de Imagen */}
                    <label>
                        URL de Imagen:
                        <input
                            type="text"
                            value={imageUrl}
                            onChange={(e) => setImageUrl(e.target.value)}
                            placeholder="Ej. /images/nombreimagen.extension"
                        />
                    </label>

                    <div className="schedules-section">
                        <h3>Horarios</h3>
                        {schedules.map((schedule, index) => (
                            <div key={index} className="schedule-item">
                                <select
                                    value={schedule.day}
                                    onChange={(e) => updateSchedule(index, "day", e.target.value)}
                                >
                                    <option value="Lunes">Lunes</option>
                                    <option value="Martes">Martes</option>
                                    <option value="Miércoles">Miércoles</option>
                                    <option value="Jueves">Jueves</option>
                                    <option value="Viernes">Viernes</option>
                                    <option value="Sábado">Sábado</option>
                                    <option value="Domingo">Domingo</option>
                                </select>
                                <input
                                    type="time"
                                    value={schedule.start_time}
                                    onChange={(e) => updateSchedule(index, "start_time", e.target.value)}
                                />
                                <input
                                    type="time"
                                    value={schedule.end_time}
                                    onChange={(e) => updateSchedule(index, "end_time", e.target.value)}
                                />
                                {schedules.length > 1 && (
                                    <button
                                        type="button"
                                        className="remove-schedule"
                                        onClick={() => removeSchedule(index)}
                                    >
                                        Eliminar
                                    </button>
                                )}
                            </div>
                        ))}
                        <button type="button" className="add-schedule" onClick={addSchedule}>
                            + Agregar Horario
                        </button>
                    </div>

                    {/* Botones */}
                    <div className="button-group">
                        <button type="submit" className="btn-submit">
                            {isEditing ? "Actualizar Actividad" : "Crear Actividad"}
                        </button>
                        {isEditing && (
                            <button type="button" className="btn-cancel" onClick={clearForm}>
                                Cancelar
                            </button>
                        )}
                    </div>
                </form>
            </div>

            {/* ===========================================================
            LISTADO DE ACTIVIDADES EXISTENTES
            =========================================================== */}
            <div className="activities-list-admin">
                {activities.length === 0 && <p>No hay actividades para mostrar.</p>}

                {activities.map((activity) => (
                    <div className="activity-card-admin" key={activity.id}>
                        {/* Header con nombre + botones de acción */}
                        <div className="activity-header">
                            <h3>{activity.name}</h3>
                            <div className="action-buttons">
                                <button className="btn-edit" onClick={() => handleEdit(activity)}>
                                    Editar
                                </button>
                                <button className="btn-delete" onClick={() => handleDelete(activity.id)}>
                                    Eliminar
                                </button>
                            </div>
                        </div>

                        {/* Imagen de la actividad */}
                        {activity.image_url && (
                            <div className="activity-image">
                                <img src={getImageUrl(activity.image_url)} alt={activity.name} />
                            </div>
                        )}

                        {/* Detalles adicionales */}
                        <p className="activity-desc">{activity.description}</p>
                        <p>
                            <strong>Capacidad:</strong> {activity.capacity}
                        </p>
                        <p>
                            <strong>Categoria:</strong> {activity.category}
                        </p>
                        <p>
                            <strong>Profesor:</strong> {activity.profesor}
                        </p>

                        {/* Horarios */}
                        <h4>Horarios:</h4>
                        <ul>
                            {Array.isArray(activity.schedules) &&
                                activities.length === 0 &&
                                activity.schedules.map((s, idx) => (
                                    <li key={idx}>
                                        {s.day}: {s.start_time} - {s.end_time}
                                    </li>
                                ))}
                        </ul>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Admn;
