namespace AntsWarForms
{
    public partial class Game : Form
    {
        public Position Highlight { get; set; }
        public GameBoard Board { get; set; }
        public Team Team { get; set; }


        public ManualResetEvent CardChooseEvent = new ManualResetEvent(false);
        public Position? CardChosen { get; private set; }


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

            BoardTable.Controls.Clear();

            for (var i = 0; i < cards.Length; i++)
            {
                for (var j = 0; j < cards[i].Length; j++)
                {
                    var card = cards[i][j];
                    var button = new Button
                    {
                        Dock = DockStyle.Fill,
                        ForeColor = card.Team == Team.Red ? Color.Red : Color.Black,
                    };

                    switch (card.Value)
                    {
                        case 0: button.Text = "Q"; break;
                        case -1: button.Text = ""; break;
                        default: button.Text = card.Value.ToString(); break;
                    }
                    button.Tag = new Position(i, j);
                    button.Click += CardClick;
                    BoardTable.Controls.Add(button, i, j);
                }
            }
        }

        public void CardClick(object? sender, EventArgs e)
        {
            if (sender is null || sender.GetType() != typeof(Button)) return;
            if (((Button)sender!).Tag is null || ((Button)sender).Tag!.GetType() != typeof(Position)) return;

            CardChosen = ((Button)sender).Tag as Position;
            CardChooseEvent.Set();
        }
    }
}
