namespace AntsWarForms
{
    public partial class Game : Form
    {
        public Position Highlight { get; set; }
        public GameBoard Board { get; set; }

        public Game()
        {
            Highlight = new Position();
            Board = new GameBoard();
            InitializeComponent();
        }

        public void Log(string text)
        {
            if (textBox1.InvokeRequired)
            {
                Action log = delegate { Log(text); };
                textBox1.Invoke(log);
                return;
            }
            textBox1.Text = text;
        }

        public void UpdateGameBoard(Card[][] cards)
        {
            if (InvokeRequired)
            {
                Action<Card[][]> update = delegate { UpdateGameBoard(cards); };
                Invoke(update, new object[] { cards });
                return;
            }

            for (var i = 0; i < cards.Length; i++)
            {
                for (var j = 0; j < cards[i].Length; j++)
                {
                    var card = cards[i][j];
                    var button = new Button
                    {
                        Dock = DockStyle.Fill,
                        //BackColor = card.Color,
                        ForeColor = card.Team == Team.Red ? Color.Red : Color.Black,
                    };

                    switch (card.Value)
                    {
                        case 0: button.Text = "Q"; break;
                        case -1: button.Text = ""; break;
                        default: button.Text = card.Value.ToString(); break;
                    }
                    //button.Click += CardClick;
                    BoardTable.Controls.Add(button, i, j);
                }
            }

            //BoardTable.Controls.Clear();
            //BoardTable.Controls.Add(BoardTable, 0, 0);
        }

        private void CardClick(object sender, EventArgs e)
        {
            var button = (Button)sender;
            //var position = (Position)button.Tag;
            //Highlight = position;
            //Log($"X: {position.X}, Y: {position.Y}");
        }
    }
}
