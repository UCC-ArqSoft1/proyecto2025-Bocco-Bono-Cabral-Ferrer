// Utilidades para manejar las llamadas a la API de inscripciones

const API_BASE_URL = 'http://localhost:8080';

// Función auxiliar para validar token
const validateToken = () => {
    const token = localStorage.getItem('token');
    if (!token) {
        throw new Error('No hay token de autenticación. Por favor, inicia sesión.');
    }
    return token;
};

// Función auxiliar para manejar errores de respuesta
const handleResponseError = async (response) => {
    const data = await response.json();

    if (response.status === 401) {
        // Token expirado o inválido
        localStorage.removeItem('token');
        localStorage.removeItem('userTypeId');
        throw new Error('Tu sesión ha expirado. Por favor, inicia sesión nuevamente.');
    }

    if (response.status === 403) {
        throw new Error('No tienes permisos para realizar esta acción.');
    }

    throw new Error(data.error || 'Error en la operación');
};

export const enrollmentApi = {
    // Verificar si un usuario está inscrito en una actividad
    checkEnrollment: async (activityId) => {
        const token = validateToken();

        const response = await fetch(`${API_BASE_URL}/enrollment/check?activity_id=${activityId}`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (!response.ok) {
            await handleResponseError(response);
        }

        const data = await response.json();
        return data.is_enrolled;
    },

    // Obtener información de capacidad de una actividad
    getActivityCapacity: async (activityId) => {
        const token = validateToken();

        const response = await fetch(`${API_BASE_URL}/enrollment/capacity?activity_id=${activityId}`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (!response.ok) {
            await handleResponseError(response);
        }

        const data = await response.json();
        return data;
    },

    // Inscribirse en una actividad
    enroll: async (activityId) => {
        const token = validateToken();

        const response = await fetch(`${API_BASE_URL}/enrollment`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({
                activity_id: activityId
            })
        });

        if (!response.ok) {
            await handleResponseError(response);
        }

        const data = await response.json();
        return data;
    },

    // Cancelar inscripción en una actividad
    cancelEnrollment: async (activityId) => {
        const token = validateToken();
        const response = await fetch(`${API_BASE_URL}/enrollment`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({
                activity_id: Number(activityId)
            })
        });

        if (!response.ok) {
            await handleResponseError(response);
        }

        const data = await response.json();
        return data;
    },

    // Obtener todas las inscripciones del usuario
    getUserEnrollments: async () => {
        const token = validateToken();

        const response = await fetch(`${API_BASE_URL}/enrollments`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (!response.ok) {
            await handleResponseError(response);
        }

        const data = await response.json();
        return data;
    }
}; 