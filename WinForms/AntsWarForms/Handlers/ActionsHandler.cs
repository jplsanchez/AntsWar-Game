using System.Text.Json;

namespace AntsWarForms.Handlers
{
    internal class ActionsHandler
    {
        private readonly Game _game;

        private readonly JsonSerializerOptions jsonOpt = new() { PropertyNameCaseInsensitive = true };

        public ActionsHandler(Game game)
        {
            _game = game;
        }

        internal void GameBoardHandle(string json)
        {
            var cards = JsonSerializer.Deserialize<DTO<Card[][]>>(json, jsonOpt)?.Data;
            if (cards is null) return;
            _game.UpdateGameBoard(cards);
        }

        internal void SingleStringHandle(string json)
        {
            var str = JsonSerializer.Deserialize<DTO<string>>(json, jsonOpt);
            if (str is null) return;

            _game.Log(str.Data);
        }

        internal string StringCollectionHandle(string json)
        {
            var strCollection = JsonSerializer.Deserialize<DTO<string[]>>(json, jsonOpt)?.Data;
            if (strCollection is null) return "";
            if (strCollection.FirstOrDefault()!.Equals("Choose a card to play"))
            {
                foreach (var message in strCollection)
                {
                    _game.Log(message);
                }
                _game.CardChooseEvent.Reset();
                _game.CardChooseEvent.WaitOne();

                if (_game.CardChosen is null) return "";
                return _game.CardChosen!.ToString();
            }

            if (strCollection.FirstOrDefault()!.Contains("Choose an action to play"))
            {
                _game.Log("Work In Progess");
                return "";
            }

            _game.Log("Work In Progess");
            return "";
        }

        internal void TeamHandle(string json)
        {
            var team = JsonSerializer.Deserialize<DTO<Team>>(json, jsonOpt);
            if (team is null) return;
            _game.Team = team.Data;
        }

        internal void PositionArrayHandle(string json)
        {
            var xyIntArray = JsonSerializer.Deserialize<DTO<List<int>>>(json, jsonOpt)?.Data;
            if (xyIntArray is null) return;

            _game.Highlight.X = xyIntArray[0];
            _game.Highlight.Y = xyIntArray[1];

        }
    }
}
