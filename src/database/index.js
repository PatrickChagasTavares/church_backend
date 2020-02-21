import Sequelize from 'sequelize';

import Children from '../app/models/Children';
import Social from '../app/models/Social';
import DoorToDoor from '../app/models/DoorToDoor';

import databseConfig from '../config/database';

const models = [Children, Social, DoorToDoor];

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
