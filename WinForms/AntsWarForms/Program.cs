using Microsoft.Extensions.Configuration;
using System.Diagnostics;

namespace AntsWarForms
{
    internal static class Program
    {
        const bool AutoRunGolangServer = true;

        /// <summary>
        ///  The main entry point for the application.
        /// </summary>
        [STAThread]
        static void Main()
        {
            if (AutoRunGolangServer) RunGolangServer();

            // To customize application configuration such as set high DPI settings or default font,
            // see https://aka.ms/applicationconfiguration.
            ApplicationConfiguration.Initialize();
            Application.Run(new Menu());
        }

        private static void RunGolangServer()
        {
            var process = new Process();
            process.StartInfo.FileName = "C:\\Users\\jpaul\\source\\repos\\Go\\ants-war-go\\antswar.exe";
            process.StartInfo.UseShellExecute = true;
            process.StartInfo.WindowStyle = ProcessWindowStyle.Minimized;
            process.Start();
        }
    }
}