import Login from "../pages/login.vue";
import Register from "../pages/register.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const authenticatedRoutes = ["/"];
// const nonAuthenticatedRoutes = ["/login", "/register"];

const routes: RouteRecordRaw[] = [
  { path: "/login", component: Login },
  { path: "/register", component: Register },
  { path: "/:pathMatch(.*)*", redirect: "/login" },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from) => {
  const token = localStorage.getItem("token");
  if (authenticatedRoutes.includes(to.path) && !token) {
    return "/login";
  }
});

export default router;
