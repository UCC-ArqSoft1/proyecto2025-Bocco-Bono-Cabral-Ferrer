import "./Activities.css";
const Activities = () => {
    const activities = [
        {
            name: "taekwondo",
            description: "Arte marcial coreana",
            schedule: [
                { day: 2, "hora-inicio": "18:30", "hora-fin": "20:00" },
                { day: 4, "hora-inicio": "18:30", "hora-fin": "20:00" }
            ]
        },
        {
            name: "zumba",
            description: "ritmos latinos",
            schedule: [
                { day: 1, "hora-inicio": "19:30", "hora-fin": "20:30" },
                { day: 3, "hora-inicio": "19:30", "hora-fin": "20:30" }
            ]
        }
    ];

    const weekDays = ["Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"];

    const isloggedin = localStorage.getItem("islogin") === "true";

    const handleEnrollment = (activityName) => {
        alert(`Inscribiendo en la actividad: ${activityName}`);
    };

    return (
        <div className="activitiescontainer">
            <h1>Actividades</h1>
            {activities.map((activity, index) => (
                <div className="activity-card" key={index}>
                    <h2>{activity.name}</h2>
                    <p>{activity.description}</p>
                    <h3>Horarios</h3>
                    <ul>
                        {activity.schedule.map((schedule, i) => (
                            <li key={i}>
                                Dia:{weekDays[schedule.day]}:
                                Hora de inicio{schedule["hora-inicio"]}Hora de fin - {schedule["hora-fin"]}
                            </li>
                        ))}
                    </ul>
                    {isloggedin && (
                        <button onClick={() => handleEnrollment(activity.name)}>Inscribir</button>
                    )
                    }
                </div>
            ))}
        </div>
    )
}
export default Activities