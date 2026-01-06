interface CloseIconProps {
    "aria-label"?: string;
    className?: string;
}

export function CloseIcon({ "aria-label": ariaLabel, className }: CloseIconProps): React.JSX.Element {
    return (
        <svg
            aria-label={ariaLabel}
            xmlns="http://www.w3.org/2000/svg"
            height="24px"
            viewBox="0 -960 960 960"
            width="24px"
            className={className}
        >
            <path d="m256-200-56-56 224-224-224-224 56-56 224 224 224-224 56 56-224 224 224 224-56 56-224-224-224 224Z" />
        </svg>
    );
}
