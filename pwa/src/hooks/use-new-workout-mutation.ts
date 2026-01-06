import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { UseMutationResult } from "@tanstack/react-query";
import { createWorkout } from "../services/workout-service";
import type { Workout } from "../dto/workout";
import { getWorkoutsQueryKey } from "./use-workouts-query";

export function useNewWorkoutMutation(): UseMutationResult<Workout, Error, string> {
    const queryClient = useQueryClient();

    const mutation = useMutation({
        mutationFn: createWorkout,
        onSuccess: async () => {
            await queryClient.invalidateQueries({ queryKey: getWorkoutsQueryKey() });
        },
    });

    return mutation;
}
