import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { useAuthStore } from "src/state";

const authenticatedRoutes = ["/"];
const nonAuthenticatedRoutes = ["/login", "/register"];

const routes: RouteRecordRaw[] = [
  { path: "/", component: () => import("../pages/Home.vue") },
  { path: "/login", component: () => import("../pages/Login.vue") },
  { path: "/register", component: () => import("../pages/Register.vue") },
  {
    path: "/forgot-password",
    component: () => import("../pages/ForgotPassword.vue"),
  },
  { path: "/teams/new", component: () => import("../pages/NewTeam.vue") },
  {
    path: "/teams/:teamId",
    component: () => import("../pages/TeamDetails.vue"),
  },
  {
    path: "/applications/new",
    component: () => import("../pages/NewApplication.vue"),
  },
  {
    path: "/applications/:applicationId",
    component: () => import("../pages/ApplicationDetails.vue"),
  },
  {
    path: "/settings",
    component: () => import("../pages/Settings.vue"),
  },
  {
    path: "/404",
    component: () => import("../pages/404.vue"),
  },
  { path: "/:pathMatch(.*)*", redirect: "/login" },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _) => {
  const { fetchAuthToken } = useAuthStore();
  const token = await fetchAuthToken();
  if (authenticatedRoutes.includes(to.path) && !token) {
    return "/login";
  } else if (nonAuthenticatedRoutes.includes(to.path) && !!token) {
    return "/";
  }
});

export default router;
