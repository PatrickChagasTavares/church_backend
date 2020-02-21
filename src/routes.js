import { Router } from 'express';
import Children from './app/models/Children';

const routes = new Router();

routes.get('/', async (req, res) => {
  const children = Children.create({
    date: '2020-02-20',
    total: 22,
    note: 'Tudo funcionando',
    create_at: new Date(),
    updated_at: new Date(),
  });

  return res.json(children);
});

export default routes;
