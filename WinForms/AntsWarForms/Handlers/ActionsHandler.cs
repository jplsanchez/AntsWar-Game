using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using System.Xml.Linq;

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
            var str = JsonSerializer.Deserialize<string>(json);
        }
        internal void StringCollectionHandle(string json)
        {
            var strCollection = JsonSerializer.Deserialize<string[]>(json);
        }
        internal void TeamHandle(string json)
        {
            Thread.Sleep(30000);
            var team = JsonSerializer.Deserialize<Team>(json);
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
