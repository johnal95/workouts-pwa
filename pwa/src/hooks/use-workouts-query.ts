import { useQuery } from "@tanstack/react-query";
import type { UseQueryResult } from "@tanstack/react-query";
import { getWorkouts } from "../services/workout-service";
import type { Workout } from "../dto/workout";

export function getWorkoutsQueryKey(): ["workout-list"] {
    return ["workout-list"];
}

export function useWorkoutsQuery(): UseQueryResult<Workout[], Error> {
    const query = useQuery({
        queryKey: getWorkoutsQueryKey(),
        queryFn: () => getWorkouts(),
    });

    return query;
}
