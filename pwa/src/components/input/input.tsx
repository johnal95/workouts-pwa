import type React from "react";

interface InputProps {
    label: string;
    name: string;
    type?: "text";
    placeholder?: string;
    value?: string | number;
    onChange?: React.ChangeEventHandler<HTMLInputElement>;
}

export function Input({ type = "text", name, label, value, placeholder, onChange }: InputProps): React.JSX.Element {
    return (
        <label className="flex w-full flex-col gap-1">
            {label}
            <input
                type={type}
                name={name}
                value={value}
                onChange={onChange}
                className="bg-surface-1 text-content rounded-xl px-4 py-2"
                placeholder={placeholder}
            />
        </label>
    );
}
