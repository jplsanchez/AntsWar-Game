using AntsWarForms.Forms;

namespace AntsWarForms
{
    public partial class Game : Form
    {
        public Position Highlight { get; set; }
        public GameBoard Board { get; set; }
        public Team Team { get; set; }


        public ManualResetEvent CardChooseEvent = new ManualResetEvent(false);
        public Position? CardChosen { get; private set; }


        public ManualResetEvent ActionChooseEvent = new ManualResetEvent(false);
        public string? ActionChosen { get; set; }

        private bool _boardLoaded = false;

        public Game()
        {
            Highlight = new Position();
            Board = new GameBoard();
            InitializeComponent();
        }

        public void Log(string text)
        {
            if (LogLabel.InvokeRequired)
            {
                Action log = delegate { Log(text); };
                LogLabel.Invoke(log);
                return;
            }

            if (text.ToLower().Contains("choose an action to play"))
            {
                var selector = new ActionSelector(this);

                selector.Show();
            }

            if (text.ToLower().StartsWith("it's your turn")) LogLabel.Text = string.Empty;

            InstructionsLbl.Text = text;
            LogLabel.Text = text + "\n" + LogLabel.Text;
        }

        public void UpdateGameBoard(Card[][] cards)
        {
            if (InvokeRequired)
            {
                Action<Card[][]> update = delegate { UpdateGameBoard(cards); };
                Invoke(update, new object[] { cards });
                return;
            }

            if (!_boardLoaded)
            {
                BoardFirstLoad(cards);
                return;
            }

            for (var i = 0; i < cards.Length; i++)
            {
                for (var j = 0; j < cards[i].Length; j++)
                {
                    var button = (Button)BoardTable.Controls.Find($"Card_{i}_{j}", true).First();
                    button.SetCardOnButton(cards[i][j]);
                }
            }
        }

        private void BoardFirstLoad(Card[][] cards)
        {
            BoardTable.Controls.Clear();

            for (var i = 0; i < cards.Length; i++)
            {
                for (var j = 0; j < cards[i].Length; j++)
                {
                    var card = cards[i][j];
                    var button = new Button
                    {
                        Dock = DockStyle.Fill,
                    };

                    button.Name = $"Card_{i}_{j}";
                    button.Tag = new Position(i, j);
                    button.Click += CardClick;
                    button.SetCardOnButton(card);

                    BoardTable.Controls.Add(button, i, j);
                }
            }
            _boardLoaded = true;
        }

        public void CardClick(object? sender, EventArgs e)
        {
            if (sender is null || sender.GetType() != typeof(Button)) return;
            if (((Button)sender!).Tag is null || ((Button)sender).Tag!.GetType() != typeof(Position)) return;

            CardChosen = ((Button)sender).Tag as Position;
            CardChooseEvent.Set();
        }
    }

    public static class ButtonExtension
    {
        public static void SetCardOnButton(this Button button, Card card)
        {
            button.ForeColor = card.Team == Team.Red ? Color.Red : Color.Black;
            switch (card.Value)
            {
                case 0: button.Text = "Q"; break;
                case -1: button.Text = ""; break;
                default: button.Text = card.Value.ToString(); break;
            }
        }
    }
}
