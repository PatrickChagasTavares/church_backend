import { Router } from 'express';

import ViagemController from './app/controllers/ViagemController';

const routes = new Router();

routes.post('/viagemMissionaria', ViagemController.store);

export default routes;
