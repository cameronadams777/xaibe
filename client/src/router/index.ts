import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const authenticatedRoutes = ["/"];
const nonAuthenticatedRoutes = ["/login", "/register"];

const routes: RouteRecordRaw[] = [
  { path: "/", component: () => import("../pages/home.vue") },
  { path: "/login", component: () => import("../pages/login.vue") },
  { path: "/register", component: () => import("../pages/register.vue") },
  { path: "/teams/new", component: () => import("../pages/new-team.vue") },
  {
    path: "/teams/:teamId",
    component: () => import("../pages/team-details.vue"),
  },
  {
    path: "/applications/new",
    component: () => import("../pages/new-application.vue"),
  },
  {
    path: "/applications/:applicationId",
    component: () => import("../pages/application-details.vue"),
  },
  {
    path: "/settings",
    component: () => import("../pages/settings.vue"),
  },
  { path: "/:pathMatch(.*)*", redirect: "/login" },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _) => {
  console.log(to);
  const token = localStorage.getItem("token");
  if (authenticatedRoutes.includes(to.path) && !token) {
    return "/login";
  } else if (nonAuthenticatedRoutes.includes(to.path) && token != null) {
    return "/";
  }
});

export default router;
