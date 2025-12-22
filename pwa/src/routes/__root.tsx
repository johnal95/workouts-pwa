import React from "react";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { ReactQueryDevtoolsPanel } from "@tanstack/react-query-devtools";
import { TanStackDevtools } from "@tanstack/react-devtools";
import { TanStackRouterDevtoolsPanel } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
    component: RootLayout,
});

function RootLayout(): React.JSX.Element {
    return (
        <React.Fragment>
            <Outlet />
            <TanStackDevtools
                config={{
                    position: "bottom-right",
                }}
                plugins={[
                    { name: "Tanstack Router", render: <TanStackRouterDevtoolsPanel /> },
                    { name: "Tanstack Query", render: <ReactQueryDevtoolsPanel /> },
                ]}
            />
        </React.Fragment>
    );
}
