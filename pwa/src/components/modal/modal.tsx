import clsx from "clsx";
import React from "react";
import { createPortal } from "react-dom";
import { m } from "../../paraglide/messages";
import { CloseIcon } from "../icons/close-icon";

interface ModalProps {
    title: string;
    isOpen: boolean;
    onClose: () => void;
    children: React.ReactNode;
    className?: string;
}

const MODAL_PORTAL_ELEMENT_ID = "modal-portal-element";

export function Modal({ title, isOpen, onClose, children, className }: ModalProps): React.JSX.Element | null {
    const modalRef = React.useRef<HTMLElement>(null);

    React.useEffect(() => {
        const existingModalElement = document.getElementById(MODAL_PORTAL_ELEMENT_ID);
        if (existingModalElement) {
            modalRef.current = existingModalElement;
        } else {
            const newModalElement = document.createElement("div");
            newModalElement.setAttribute("id", MODAL_PORTAL_ELEMENT_ID);
            document.body.appendChild(newModalElement);
            modalRef.current = newModalElement;
        }
    }, []);

    if (!modalRef.current || !isOpen) {
        return null;
    }

    return createPortal(
        <div className="fixed inset-0 z-50 flex items-center justify-center" role="dialog" aria-modal="true">
            <div onClick={onClose} className="absolute inset-0 bg-black/50" aria-hidden="true" />
            <div className={clsx("bg-surface-0 relative h-full w-full sm:h-auto sm:w-auto sm:rounded-xl", className)}>
                <header className="border-primary-dim/20 grid grid-cols-[1fr_2fr_1fr] border-b p-4">
                    <div>
                        <button type="button" className="cursor-pointer" onClick={onClose}>
                            <CloseIcon aria-label={m.modal_close_icon_aria_label()} className="fill-content" />
                        </button>
                    </div>
                    <h1 className="justify-self-center font-bold">{title}</h1>
                </header>
                <div className="p-4">{children}</div>
            </div>
        </div>,
        modalRef.current,
    );
}
