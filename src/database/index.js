import Sequelize from 'sequelize';

import Children from '../app/models/Children';

import databseConfig from '../config/database';

const models = [Children];

class Database {
  constructor() {
    this.init();
  }

  init() {
    this.connection = new Sequelize(databseConfig);

    models.map(model => model.init(this.connection));
  }
}

export default new Database();
