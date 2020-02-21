module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('door_to_doors', {
      id: {
        type: Sequelize.INTEGER,
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
      },
      date: {
        type: Sequelize.DATE,
        allowNull: false,
      },
      name_tribe: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      name_person: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      address: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      type: {
        type: Sequelize.STRING(20),
        allowNull: false,
      },
      age: {
        type: Sequelize.STRING(20),
        allowNull: false,
      },
      bible: {
        type: Sequelize.STRING(20),
      },
      evangelical: {
        type: Sequelize.STRING(20),
      },
      contact: {
        type: Sequelize.STRING(20),
      },
      frequents_church: {
        type: Sequelize.STRING(20),
      },
      cult_home: {
        type: Sequelize.STRING(20),
      },
      study_bible: {
        type: Sequelize.STRING(20),
      },
      study_confimation: {
        type: Sequelize.STRING(20),
      },
      prayer_request: {
        type: Sequelize.STRING(20),
      },
      reconciled: {
        type: Sequelize.STRING(20),
      },
      visit: {
        type: Sequelize.STRING(20),
      },
      accept_christ: {
        type: Sequelize.STRING(20),
      },
      note: {
        type: Sequelize.STRING,
      },
      created_at: {
        type: Sequelize.DATE,
        allowNull: false,
      },
      updated_at: {
        type: Sequelize.DATE,
        allowNull: false,
      },
    });
  },

  down: queryInterface => {
    return queryInterface.dropTable('door_to_doors');
  },
};
