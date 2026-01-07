import React from "react";
import { createFileRoute } from "@tanstack/react-router";
import { TopHeader } from "../components/header/top-header";
import { NavigationCard } from "../components/cards/navigation-card";
import { AddCard } from "../components/cards/add-card";
import { Modal } from "../components/modal/modal";
import { Button } from "../components/buttons/button";
import { useWorkoutsQuery } from "../hooks/use-workouts-query";
import { useNewWorkoutMutation } from "../hooks/use-new-workout-mutation";
import { m } from "../paraglide/messages";
import { Input } from "../components/input/input";

export const Route = createFileRoute("/workouts/")({
    component: RouteComponent,
});

interface AddWorkoutModalContentProps {
    closeModal: () => void;
}

function AddWorkoutModalContent({ closeModal }: AddWorkoutModalContentProps) {
    const [workoutName, setWorkoutName] = React.useState("");

    const newWorkoutMutation = useNewWorkoutMutation();

    return (
        <form
            className="flex flex-col gap-4"
            onSubmit={async (e) => {
                e.preventDefault();
                await newWorkoutMutation.mutateAsync(workoutName);
                closeModal();
            }}
        >
            <Input
                label={m.workouts_add_new_name_input_label()}
                name="workout-name"
                value={workoutName}
                placeholder={m.workouts_add_new_name_input_placeholder()}
                onChange={(e) => {
                    setWorkoutName(e.target.value);
                }}
            />
            <Button type="submit">{m.workouts_add_new_submit_cta()}</Button>
        </form>
    );
}

function RouteComponent() {
    const [isNewWorkoutModalOpen, setIsNewWorkoutModalOpen] = React.useState(false);

    const workoutsQuery = useWorkoutsQuery();

    if (workoutsQuery.isLoading) {
        // TODO : Improve loading UI
        return <p>{"Loading..."}</p>;
    }

    if (!workoutsQuery.isSuccess) {
        // TODO : Improve error UI
        return (
            <div>
                <p>{"Error!"}</p>
                {workoutsQuery.error?.message && <p>{`Detail: ${workoutsQuery.error.message}`}</p>}
            </div>
        );
    }

    return (
        <div>
            <TopHeader heading={m.workouts_title()} />
            <div className="flex flex-col gap-4 p-4">
                {workoutsQuery.data.map((w) => (
                    <NavigationCard
                        key={w.id}
                        to="/workouts/$workoutId"
                        params={{ workoutId: w.id }}
                        icon="eventList"
                        text={w.name}
                        textAs="h2"
                    />
                ))}
                <AddCard
                    text={m.workouts_card_add_new_cta()}
                    onClick={() => {
                        setIsNewWorkoutModalOpen(true);
                    }}
                />
            </div>
            <Modal
                isOpen={isNewWorkoutModalOpen}
                onClose={() => {
                    setIsNewWorkoutModalOpen(false);
                }}
                title={m.workouts_add_new_title()}
            >
                <AddWorkoutModalContent
                    closeModal={() => {
                        setIsNewWorkoutModalOpen(false);
                    }}
                />
            </Modal>
        </div>
    );
}
