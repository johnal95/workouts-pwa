import { PlusIcon } from "../icons/plus-icon";

interface AddCardProps {
    text: string;
    onClick: React.MouseEventHandler<HTMLButtonElement>;
}

export function AddCard({ text, onClick }: AddCardProps) {
    return (
        <button
            className="border-content-secondary text-content-secondary flex cursor-pointer items-center justify-center gap-2 rounded-xl border border-dashed p-4"
            onClick={onClick}
        >
            <PlusIcon className="fill-content-secondary" />
            {text}
        </button>
    );
}
