import React, { useState, useEffect } from "react";
import "../Styles/activitiesAdmin.css";

const Admn = () => {
    // ------------------------------
    // 1) ESTADOS
    // ------------------------------
    const [activities, setActivities] = useState([]);

    // Para controlar si estamos editando o creando
    const [isEditing, setIsEditing] = useState(false);
    const [editId, setEditId] = useState(null);

    // Campos del formulario
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [capacity, setCapacity] = useState(0);
    const [category, setCategory] = useState("");
    const [profesor, setProfesor] = useState("");

    // Para “Horarios” usamos un textarea que contenga un JSON de ActivitySchedule[]
    const [scheduleText, setScheduleText] = useState(
        `[
  {
    "day": "Lunes",
    "start_time": "19:00",
    "end_time": "20:00"
  }
]`
    );

    // Si tu API requiere token, lo recogemos de localStorage (o de donde lo tengas)
    const token = localStorage.getItem("token") || "";

    // ------------------------------
    // 2) FETCH DE ACTIVIDADES (GET /activities)
    // ------------------------------
    const fetchActivities = async () => {
        try {
            const resp = await fetch("/activities", {
                headers: {
                    "Content-Type": "application/json",
                    ...(token ? { Authorization: `Bearer ${token}` } : {}),
                },
            });
            if (!resp.ok) throw new Error("Error al obtener actividades");

            const data = await resp.json();
            // Suponemos que el backend devuelve { activities: [...] } o un array directamente
            setActivities(data.activities || data);
        } catch (error) {
            console.error("Error al cargar actividades:", error);
            alert("No se pudieron cargar las actividades. Revisa la consola.");
        }
    };

    useEffect(() => {
        fetchActivities();
        // eslint-disable-next-line react-hooks/exhaustive-deps
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
        setScheduleText(
            `[
  {
    "day": "Lunes",
    "start_time": "19:00",
    "end_time": "20:00"
  }
]`
        );
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

        // ----- Parsear el JSON de 'scheduleText' -----
        let schedulesArr;
        try {
            schedulesArr = JSON.parse(scheduleText);
            if (!Array.isArray(schedulesArr)) {
                throw new Error("Debe ser un arreglo de objetos JSON");
            }
            // Verificar que cada objeto tenga las claves esperadas
            schedulesArr.forEach((s, idx) => {
                if (
                    typeof s.day !== "string" ||
                    typeof s.start_time !== "string" ||
                    typeof s.end_time !== "string"
                ) {
                    throw new Error(
                        `El elemento en la posición ${idx} no tiene las propiedades "day", "start_time" y "end_time" como strings.`
                    );
                }
            });
        } catch (err) {
            alert(
                "El campo de Horarios debe ser un JSON válido.\n\nEjemplo de formato:\n" +
                `[
  {
    "day": "Lunes",
    "start_time": "19:00",
    "end_time": "20:00"
  },
  {
    "day": "Miércoles",
    "start_time": "18:30",
    "end_time": "20:00"
  }
]`
            );
            return;
        }

        // ----- Construir el body según el dominio -----
        const body = {
            name: name.trim(),
            description: description.trim(),
            capacity: parseInt(capacity, 10),
            category: category.trim(),
            profesor: profesor.trim(),
            schedules: schedulesArr, // Aquí va el array de ActivitySchedule
        };

        try {
            let resp;
            if (isEditing && editId) {
                // --> UPDATE (PUT /activities/:id)
                resp = await fetch(`/activities/${editId}`, {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        ...(token ? { Authorization: `Bearer ${token}` } : {}),
                    },
                    body: JSON.stringify(body),
                });
                if (!resp.ok) throw new Error("Error al actualizar");
                alert("Actividad actualizada correctamente.");
            } else {
                // --> CREATE (POST /activities)
                resp = await fetch("/activities", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        ...(token ? { Authorization: `Bearer ${token}` } : {}),
                    },
                    body: JSON.stringify(body),
                });
                if (!resp.ok) throw new Error("Error al crear");
                alert("Actividad creada correctamente.");
            }

            clearForm();
            fetchActivities();
        } catch (err) {
            console.error(err);
            alert("Ocurrió un problema con el servidor. Revisa la consola.");
        }
    };

    // ------------------------------
    // 5) ELIMINAR ACTIVIDAD (DELETE /activities/:id)
    // ------------------------------
    const handleDelete = async (id) => {
        if (!window.confirm("¿Estás seguro que deseas eliminar esta actividad?")) return;
        try {
            const resp = await fetch(`/activities/${id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    ...(token ? { Authorization: `Bearer ${token}` } : {}),
                },
            });
            if (!resp.ok) throw new Error("Error al eliminar");
            alert("Actividad eliminada.");
            fetchActivities();
        } catch (err) {
            console.error(err);
            alert("No se pudo eliminar la actividad. Revisa la consola.");
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
        // Convertimos schedules a JSON con identación
        setScheduleText(JSON.stringify(activity.schedules, null, 2));

        // Opcional: desplazarse al formulario
        document.getElementById("admin-form")?.scrollIntoView({ behavior: "smooth" });
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

                    {/* 6) Horarios (JSON) */}
                    <label>
                        Horarios (JSON de ActivitySchedule[]):
                        <textarea
                            className="schedule-textarea"
                            rows={6}
                            value={scheduleText}
                            onChange={(e) => setScheduleText(e.target.value)}
                            placeholder={`Ejemplo de JSON para "schedules":\n[
  {
    "day": "Lunes",
    "start_time": "19:00",
    "end_time": "20:00"
  },
  {
    "day": "Miércoles",
    "start_time": "18:30",
    "end_time": "20:00"
  }
]`}
                        />
                    </label>

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
                                activity.schedules.map((s, idx) => (
                                    <li key={idx}>
                                        Dia: <strong>{s.day}</strong> – Inicio: {s.start_time} – Fin: {s.end_time}
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
