namespace AntsWarForms
{
    public record DTO<T>(string Name, T Data);

    public class GameBoard
    {
        public Card[][]? Cards { get; set; }
    }

    public class Card
    {
        public int Value { get; set; }
        public Team Team { get; set; }
    }

    public enum Team
    {
        Red,
        Blue
    }

    public class Position
    {
        public int X { get; set; }
        public int Y { get; set; }

        public Position()
        {
            X = 0;
            Y = 0;
        }
    }
}
