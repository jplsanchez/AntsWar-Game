namespace AntsWarForms.Forms
{
    partial class ActionSelector
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
            AttackBtn = new Button();
            SwapBtn = new Button();
            MoveBtn = new Button();
            MarchBtn = new Button();
            SuspendLayout();
            // 
            // AttackBtn
            // 
            AttackBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            AttackBtn.Location = new Point(225, 78);
            AttackBtn.Name = "AttackBtn";
            AttackBtn.Size = new Size(207, 60);
            AttackBtn.TabIndex = 10;
            AttackBtn.Text = "Atacar";
            AttackBtn.UseVisualStyleBackColor = true;
            AttackBtn.Click += AttackBtn_Click;
            // 
            // SwapBtn
            // 
            SwapBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            SwapBtn.Location = new Point(12, 78);
            SwapBtn.Name = "SwapBtn";
            SwapBtn.Size = new Size(207, 60);
            SwapBtn.TabIndex = 9;
            SwapBtn.Text = "Trocar";
            SwapBtn.UseVisualStyleBackColor = true;
            SwapBtn.Click += SwapBtn_Click;
            // 
            // MoveBtn
            // 
            MoveBtn.Font = new Font("Segoe UI", 11F, FontStyle.Regular, GraphicsUnit.Point);
            MoveBtn.Location = new Point(225, 12);
            MoveBtn.Name = "MoveBtn";
            MoveBtn.Size = new Size(207, 60);
            MoveBtn.TabIndex = 8;
            MoveBtn.Text = "Mover";
            MoveBtn.UseVisualStyleBackColor = true;
            MoveBtn.Click += MoveBtn_Click;
            // 
            // MarchBtn
            // 
            MarchBtn.Font = new Font("Segoe UI", 11.25F, FontStyle.Regular, GraphicsUnit.Point);
            MarchBtn.Location = new Point(12, 12);
            MarchBtn.Name = "MarchBtn";
            MarchBtn.Size = new Size(207, 60);
            MarchBtn.TabIndex = 7;
            MarchBtn.Text = "Marchar";
            MarchBtn.UseVisualStyleBackColor = true;
            MarchBtn.Click += MarchBtn_Click;
            // 
            // ActionSelector
            // 
            AutoScaleDimensions = new SizeF(7F, 15F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(445, 147);
            Controls.Add(AttackBtn);
            Controls.Add(SwapBtn);
            Controls.Add(MoveBtn);
            Controls.Add(MarchBtn);
            FormBorderStyle = FormBorderStyle.FixedDialog;
            Name = "ActionSelector";
            StartPosition = FormStartPosition.CenterScreen;
            Text = "ActionSelector";
            ResumeLayout(false);
        }

        #endregion

        private Button AttackBtn;
        private Button SwapBtn;
        private Button MoveBtn;
        private Button MarchBtn;
    }
}