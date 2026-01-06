import Type from "typebox";
import { Workout } from "../dto/workout";
import { fetchJSON } from "../utilities/fetch-util";

export async function getWorkout(workoutId: string): Promise<Workout> {
    return fetchJSON(Workout, `/api/v1/workouts/${workoutId}`);
}

export async function getWorkouts(): Promise<Workout[]> {
    return fetchJSON(Type.Array(Workout), "/api/v1/workouts");
}

export async function createWorkout(workoutName: string): Promise<Workout> {
    return fetchJSON(Workout, "/api/v1/workouts", {
        method: "POST",
        body: JSON.stringify({ name: workoutName }),
        headers: { "content-type": "application/json" },
    });
}
