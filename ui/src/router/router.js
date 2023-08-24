// router.js
import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '../pages/HomePage'; // Import your components
import ConnectorsPage from '../pages/ConnectorsPage';
import AnalyzersPage from '../pages/AnalyzersPage';
import IntegratorsPage from '../pages/IntegratorsPage';
import DataPage from '../pages/DataPage'
const routes = [
  {
    path: '/',
    component: HomePage,
  },
  {
    path: '/connectors',
    component: ConnectorsPage
  },
  {
    path: '/analyzers',
    component: AnalyzersPage
  },
  {
    path: '/integrators',
    component: IntegratorsPage
  },
  {
    path: '/data',
    component: DataPage
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
