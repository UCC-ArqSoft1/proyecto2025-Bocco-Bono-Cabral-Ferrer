import { useState, useEffect } from 'react';
import { useAuth } from './useAuth';
import { enrollmentApi } from '../utils/enrollmentApi.jsx';

export const useEnrollment = (activityId) => {
    const [isEnrolled, setIsEnrolled] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    // Verificar estado de inscripción al cargar
    useEffect(() => {
        const checkEnrollmentStatus = async () => {
            if (!isAuthenticated || isAdmin || !activityId) return;

            try {
                const enrolled = await enrollmentApi.checkEnrollment(activityId);
                setIsEnrolled(enrolled);
            } catch (err) {
                console.error('Error checking enrollment status:', err);
                setError(err.message);
            }
        };

        checkEnrollmentStatus();
    }, [activityId, isAuthenticated, isAdmin]);

    // Función para inscribirse
    const enroll = async () => {
        if (!isAuthenticated) {
            throw new Error('Debes estar autenticado para inscribirte');
        }

        setIsLoading(true);
        setError(null);

        try {
            await enrollmentApi.enroll(activityId);
            setIsEnrolled(true);
            return { success: true };
        } catch (err) {
            setError(err.message);
            throw err;
        } finally {
            setIsLoading(false);
        }
    };

    // Función para cancelar inscripción
    const cancelEnrollment = async () => {
        if (!isAuthenticated) {
            throw new Error('Debes estar autenticado para cancelar la inscripción');
        }

        setIsLoading(true);
        setError(null);

        try {
            await enrollmentApi.cancelEnrollment(activityId);
            setIsEnrolled(false);
            return { success: true };
        } catch (err) {
            setError(err.message);
            throw err;
        } finally {
            setIsLoading(false);
        }
    };

    // Función para alternar inscripción (inscribirse si no está inscrito, cancelar si está inscrito)
    const toggleEnrollment = async () => {
        if (isEnrolled) {
            return await cancelEnrollment();
        } else {
            return await enroll();
        }
    };

    return {
        isEnrolled,
        isLoading,
        error,
        enroll,
        cancelEnrollment,
        toggleEnrollment
    };
}; 