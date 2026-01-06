import { Link } from "@tanstack/react-router";
import type { LinkComponentProps } from "@tanstack/react-router";
import { EventListIcon } from "../icons/event-list-icon";
import { DumbbellIcon } from "../icons/dumbell-icon";

const icons = {
    eventList: EventListIcon,
    dumbell: DumbbellIcon,
} as const;

interface NavigationCardProps extends LinkComponentProps {
    text: string;
    textAs?: `h${1 | 2 | 3 | 4 | 5 | 6}` | "p";
    icon?: keyof typeof icons;
}

export function NavigationCard({ text, icon, textAs: Text = "p", ...props }: NavigationCardProps): React.JSX.Element {
    const Icon = icon ? icons[icon] : null;

    return (
        <Link {...props} className="bg-surface flex items-center gap-2 rounded-xl p-4">
            {Icon && (
                <div className="bg-surface-0 rounded-xl p-2">
                    <Icon className="fill-primary" />
                </div>
            )}
            <Text>{text}</Text>
        </Link>
    );
}
