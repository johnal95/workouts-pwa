import clsx from "clsx";

interface ButtonProps extends React.DetailedHTMLProps<
    React.ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
> {
    intent?: "primary" | "secondary" | "ghost";
}

export function Button({ children, intent = "primary", className, ...props }: ButtonProps) {
    return (
        <button
            {...props}
            className={clsx(
                "text-surface-0 cursor-pointer rounded-xl px-4 py-2",
                {
                    "bg-primary": intent === "primary",
                },
                className,
            )}
        >
            {children}
        </button>
    );
}
