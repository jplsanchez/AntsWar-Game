namespace AntsWarForms.Forms
{
    public partial class ActionSelector : Form
    {
        private Game _game;

        public ActionSelector(Game game)
        {
            InitializeComponent();
            _game = game;
        }

        private void MarchBtn_Click(object sender, EventArgs e)
        {
            _game.ActionChosen = "march";
            _game.ActionChooseEvent.Set();
            this.Close();
        }

        private void MoveBtn_Click(object sender, EventArgs e)
        {
            _game.ActionChosen = "move";
            _game.ActionChooseEvent.Set();
            this.Close();
        }

        private void SwapBtn_Click(object sender, EventArgs e)
        {
            _game.ActionChosen = "swap";
            _game.ActionChooseEvent.Set();
            this.Close();
        }

        private void AttackBtn_Click(object sender, EventArgs e)
        {
            _game.ActionChosen = "attack";
            _game.ActionChooseEvent.Set();
            this.Close();
        }
    }
}
