import DashboardLayout from "../pages/Layout/DashboardLayout.vue";
import Producers from "../pages/Producers.vue";

const routes = [
    {
        path: "/",
        component: DashboardLayout,
        redirect: "/producers",
        children: [
            {
                path: "producers",
                name: "Producers",
                component: Producers,
            },
            {
                path: "consumers",
                name: "Consumers",
                component: Producers,
            },
        ],
    },
];

export default routes;
