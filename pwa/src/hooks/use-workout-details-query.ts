import { useQuery } from "@tanstack/react-query";
import type { UseQueryResult } from "@tanstack/react-query";
import { getWorkout } from "../services/workout-service";
import type { Workout } from "../dto/workout";

export function getWorkoutDetailsQueryKey(workoutId: string): ["workout-details", { id: string }] {
    return ["workout-details", { id: workoutId }];
}

export function useWorkoutDetailsQuery(workoutId: string): UseQueryResult<Workout, Error> {
    const query = useQuery({
        queryKey: getWorkoutDetailsQueryKey(workoutId),
        queryFn: () => getWorkout(workoutId),
    });

    return query;
}
