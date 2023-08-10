using System.Diagnostics;
using System.Xml.Linq;
using AntsWarForms.Handlers;

namespace AntsWarForms
{
    public partial class Menu : Form
    {
        private GameHandler _handler;
        private Game _game;

        public Menu()
        {
            InitializeComponent();
            _game = new();
            _handler = new(_game);
        }

        private async void StartGame(object sender, EventArgs e)
        {
            await _handler.TryToConnect();

            _game.Show();
            this.Hide();

            await Task.Run(async () =>
            {
                while (true)
                {
                    await _handler.ReceiveData();
                }
            });
        }
    }
}
