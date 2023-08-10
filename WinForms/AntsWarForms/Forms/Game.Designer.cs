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
            textBox1 = new TextBox();
            Connect = new Button();
            BoardTable = new TableLayoutPanel();
            SuspendLayout();
            // 
            // textBox1
            // 
            textBox1.Location = new Point(496, 410);
            textBox1.Multiline = true;
            textBox1.Name = "textBox1";
            textBox1.Size = new Size(211, 28);
            textBox1.TabIndex = 0;
            // 
            // Connect
            // 
            Connect.Location = new Point(713, 415);
            Connect.Name = "Connect";
            Connect.Size = new Size(75, 23);
            Connect.TabIndex = 1;
            Connect.Text = "Connect";
            Connect.UseVisualStyleBackColor = true;
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
            BoardTable.Location = new Point(124, 12);
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
            BoardTable.Size = new Size(269, 398);
            BoardTable.TabIndex = 2;
            // 
            // Game
            // 
            AutoScaleDimensions = new SizeF(7F, 15F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(800, 450);
            Controls.Add(BoardTable);
            Controls.Add(Connect);
            Controls.Add(textBox1);
            Name = "Game";
            Text = "Game";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private TextBox textBox1;
        private Button Connect;
        private TableLayoutPanel BoardTable;
    }
}