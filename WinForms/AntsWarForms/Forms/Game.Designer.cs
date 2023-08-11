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
            MarchBtn = new Button();
            MoveBtn = new Button();
            SwapBtn = new Button();
            AttackBtn = new Button();
            LogLabel = new Label();
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
            BoardTable.Size = new Size(474, 612);
            BoardTable.TabIndex = 2;
            // 
            // MarchBtn
            // 
            MarchBtn.Font = new Font("Segoe UI", 11.25F, FontStyle.Regular, GraphicsUnit.Point);
            MarchBtn.Location = new Point(687, 83);
            MarchBtn.Name = "MarchBtn";
            MarchBtn.Size = new Size(207, 60);
            MarchBtn.TabIndex = 3;
            MarchBtn.Text = "Marchar";
            MarchBtn.UseVisualStyleBackColor = true;
            MarchBtn.Click += MarchBtn_Click;
            // 
            // MoveBtn
            // 
            MoveBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            MoveBtn.Location = new Point(687, 182);
            MoveBtn.Name = "MoveBtn";
            MoveBtn.Size = new Size(207, 60);
            MoveBtn.TabIndex = 4;
            MoveBtn.Text = "Mover";
            MoveBtn.UseVisualStyleBackColor = true;
            MoveBtn.Click += MoveBtn_Click;
            // 
            // SwapBtn
            // 
            SwapBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            SwapBtn.Location = new Point(687, 283);
            SwapBtn.Name = "SwapBtn";
            SwapBtn.Size = new Size(207, 60);
            SwapBtn.TabIndex = 5;
            SwapBtn.Text = "Trocar";
            SwapBtn.UseVisualStyleBackColor = true;
            SwapBtn.Click += SwapBtn_Click;
            // 
            // AttackBtn
            // 
            AttackBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            AttackBtn.Location = new Point(687, 398);
            AttackBtn.Name = "AttackBtn";
            AttackBtn.Size = new Size(207, 60);
            AttackBtn.TabIndex = 6;
            AttackBtn.Text = "Atacar";
            AttackBtn.UseVisualStyleBackColor = true;
            AttackBtn.Click += AttackBtn_Click;
            // 
            // LogLabel
            // 
            LogLabel.AutoSize = true;
            LogLabel.Font = new Font("Segoe UI", 11.25F, FontStyle.Regular, GraphicsUnit.Point);
            LogLabel.Location = new Point(687, 503);
            LogLabel.Name = "LogLabel";
            LogLabel.Size = new Size(40, 20);
            LogLabel.TabIndex = 7;
            LogLabel.Text = "Logs";
            // 
            // Game
            // 
            AutoScaleDimensions = new SizeF(7F, 15F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(1005, 664);
            Controls.Add(LogLabel);
            Controls.Add(AttackBtn);
            Controls.Add(SwapBtn);
            Controls.Add(MoveBtn);
            Controls.Add(MarchBtn);
            Controls.Add(BoardTable);
            Name = "Game";
            StartPosition = FormStartPosition.CenterParent;
            Text = "Game";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion
        private TableLayoutPanel BoardTable;
        private Button MarchBtn;
        private Button MoveBtn;
        private Button SwapBtn;
        private Button AttackBtn;
        private Label LogLabel;
    }
}