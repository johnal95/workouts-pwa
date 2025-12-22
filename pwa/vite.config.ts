import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
    build: {
        outDir: "../static/dist/",
    },
    server: {
        proxy: {
            "/api": { target: "http://localhost:8080" },
        },
    },
    plugins: [
        tailwindcss(),
        tanstackRouter({
            target: "react",
            autoCodeSplitting: true,
            generatedRouteTree: "./src/route-tree.gen.tsx",
        }),
        react({
            babel: {
                plugins: [["babel-plugin-react-compiler"]],
            },
        }),
    ],
});
