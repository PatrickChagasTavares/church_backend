module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('socials', {
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
      medical: {
        type: Sequelize.STRING(20),
      },
      optician: {
        type: Sequelize.STRING(20),
      },
      pressure: {
        type: Sequelize.STRING(20),
      },
      glucose: {
        type: Sequelize.STRING(20),
      },
      aesthetics: {
        type: Sequelize.STRING(20),
      },
      cutting_hair: {
        type: Sequelize.STRING(20),
      },
      hairstyle: {
        type: Sequelize.STRING(20),
      },
      Nail: {
        type: Sequelize.STRING(20),
      },
      Eyebrow: {
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
    return queryInterface.dropTable('socials');
  },
};
