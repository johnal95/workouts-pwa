import { createFileRoute } from "@tanstack/react-router";
import { useWorkoutDetailsQuery } from "../hooks/use-workout-details-query";
import { TopHeader } from "../components/header/top-header";

export const Route = createFileRoute("/workouts/$workoutId/")({
    component: RouteComponent,
});

function RouteComponent() {
    const { workoutId } = Route.useParams();

    const workoutDetailsQuery = useWorkoutDetailsQuery(workoutId);

    if (workoutDetailsQuery.isLoading) {
        // TODO : Improve loading UI
        return <p>{"Loading..."}</p>;
    }

    if (!workoutDetailsQuery.isSuccess) {
        // TODO : Improve error UI
        return (
            <div>
                <p>{"Error!"}</p>
                {workoutDetailsQuery.error?.message && <p>{`Detail: ${workoutDetailsQuery.error.message}`}</p>}
            </div>
        );
    }

    return (
        <div>
            <TopHeader
                heading={workoutDetailsQuery.data.name}
                backLink={{
                    to: "/workouts",
                }}
            />
            <div>{"建設中"}</div>
        </div>
    );
}
