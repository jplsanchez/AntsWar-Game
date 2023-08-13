namespace AntsWarForms
{
    partial class Game
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            BoardTable = new TableLayoutPanel();
            LogLabel = new Label();
            InstructionsLbl = new Label();
            SuspendLayout();
            // 
            // BoardTable
            // 
            BoardTable.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;
            BoardTable.ColumnCount = 5;
            BoardTable.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 20F));
            BoardTable.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 20F));
            BoardTable.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 20F));
            BoardTable.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 20F));
            BoardTable.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 20F));
            BoardTable.Font = new Font("Segoe UI", 18F, FontStyle.Bold, GraphicsUnit.Point);
            BoardTable.Location = new Point(12, 82);
            BoardTable.Name = "BoardTable";
            BoardTable.RowCount = 8;
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.RowStyles.Add(new RowStyle(SizeType.Percent, 12.5F));
            BoardTable.Size = new Size(437, 570);
            BoardTable.TabIndex = 2;
            // 
            // LogLabel
            // 
            LogLabel.AutoSize = true;
            LogLabel.Font = new Font("Segoe UI", 11.25F, FontStyle.Regular, GraphicsUnit.Point);
            LogLabel.Location = new Point(455, 9);
            LogLabel.Name = "LogLabel";
            LogLabel.Size = new Size(166, 20);
            LogLabel.TabIndex = 7;
            LogLabel.Text = "lt's your turn! Team:Blue";
            // 
            // InstructionsLbl
            // 
            InstructionsLbl.AutoSize = true;
            InstructionsLbl.Font = new Font("Segoe UI", 20.25F, FontStyle.Bold, GraphicsUnit.Point);
            InstructionsLbl.Location = new Point(12, 9);
            InstructionsLbl.Name = "InstructionsLbl";
            InstructionsLbl.Size = new Size(157, 37);
            InstructionsLbl.TabIndex = 8;
            InstructionsLbl.Text = "Intructions";
            // 
            // Game
            // 
            AutoScaleDimensions = new SizeF(7F, 15F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(622, 664);
            Controls.Add(InstructionsLbl);
            Controls.Add(LogLabel);
            Controls.Add(BoardTable);
            Name = "Game";
            StartPosition = FormStartPosition.CenterScreen;
            Text = "Game";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion
        private TableLayoutPanel BoardTable;
        private Label LogLabel;
        private Label InstructionsLbl;
    }
}