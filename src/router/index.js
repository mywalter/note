import Vue from 'vue'
import VueRouter from 'vue-router'
import layout from '../views/layout'
import lists from "../views/lists";
import editor from "../views/editor";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: layout,
    children: [{
      path: '/:id',
      name: 'lists',
      component: lists,
      children: [{
        path: '/:id/:uuid',
        name: 'editor',
        component: editor
      }]
    }]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
