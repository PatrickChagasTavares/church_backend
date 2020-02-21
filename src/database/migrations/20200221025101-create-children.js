module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('children', {
      id: {
        type: Sequelize.INTEGER,
        allowNull: false,
        autoImcrement: true,
        primaryKey: true,
      },
      date: { type: Sequelize.DATE, allowNull: false },
      total: { type: Sequelize.INTEGER, allowNull: false },
      note: { type: Sequelize.STRING },
      create_at: {
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
    return queryInterface.dropTable('children');
  },
};
