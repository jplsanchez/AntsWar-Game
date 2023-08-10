using System.Net;
using System.Net.Sockets;
using System.Text;

namespace AntsWarForms.Handlers
{
    public class GameHandler
    {
        private Socket? _handler;
        private Game _game;
        private readonly ActionsHandler _actionsHandler;

        public GameHandler(Game game)
        {
            _game = game;
            _actionsHandler = new(_game);
        }

        public async Task TryToConnect()
        {
            var ipAddress = Dns.GetHostEntry("localhost").AddressList[0];
            IPEndPoint ipEndPoint = new(ipAddress, 8080);
            try
            {
                _handler = new(
                    ipEndPoint.AddressFamily,
                    SocketType.Stream,
                    ProtocolType.Tcp);

                await _handler.ConnectAsync(ipEndPoint);
            }
            catch (Exception)
            {
                Thread.Sleep(1000);
                await TryToConnect();
            }
        }

        public async Task ReceiveData()
        {
            if (_handler is null) await TryToConnect();

            var buffer = new byte[2048];
            var received = await _handler!.ReceiveAsync(buffer, SocketFlags.None);
            if (received == 0) return;
            var response = Encoding.UTF8.GetString(buffer, 0, received);

            HandleSocketResponse(response);
        }


        private void HandleSocketResponse(string response)
        {
            string answer = "OK";

            if (response.Contains("\"name\":\"GameBoard\"")) _actionsHandler.GameBoardHandle(response);
            else if (response.Contains("\"name\":\"SingleString\"")) _actionsHandler.SingleStringHandle(response);
            else if (response.Contains("\"name\":\"StringCollection\"")) answer = _actionsHandler.StringCollectionHandle(response);
            else if (response.Contains("\"name\":\"Team\"")) _actionsHandler.TeamHandle(response);
            else if (response.Contains("\"name\":\"PositionArray\"")) _actionsHandler.PositionArrayHandle(response);

            _handler!.Send(Encoding.UTF8.GetBytes(answer));
        }
    }
}