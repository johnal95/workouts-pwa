import { Link } from "@tanstack/react-router";
import type { LinkOptions } from "@tanstack/react-router";
import { ArrowAltIcon } from "../icons/arrow-alt-icon";

interface HeaderProps {
    heading: string;
    backLink?: Pick<LinkOptions, "to" | "params">;
}

function TopHeader({ backLink, heading }: HeaderProps): React.JSX.Element {
    return (
        <div className="border-primary-dim/20 flex justify-between border-b p-4">
            <div className="w-4">
                {backLink && (
                    <Link to={backLink.to} params={backLink.params}>
                        <ArrowAltIcon orientation="left" className="fill-content" />
                    </Link>
                )}
            </div>
            <h1 className="font-bold">{heading}</h1>
            <div className="w-4"></div>
        </div>
    );
}

export { TopHeader };
