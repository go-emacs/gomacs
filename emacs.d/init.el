(menu-bar-mode 0)
(transient-mark-mode t)
(show-paren-mode t)
(setq visible-bell t)
(setq ring-bell-function 'ignore)
(setq-default indent-tabs-mode t)
(setq make-backup-files nil)
(setq auto-save-default nil)
(setq uniquify-buffer-name-style 'post-forward-angle-brackets)

(define-key global-map "\C-h" 'delete-backward-char)
(global-set-key "\M-g" 'goto-line)

(require 'go-mode-load)
(require 'golint)
;(add-hook 'go-mode-hook (lambda () (go-oracle-mode t) ))

(add-hook 'go-mode-hook 
          (lambda()
             ;; tab size is 4
             (setq tab-width 4)
	     ;; C-c c compile
             (setq compile-command "go test -v")
	     (define-key go-mode-map "\C-cc" 'compile)
	     ;; C-c C-c 
	     (define-key go-mode-map "\C-c\C-c" 'comment-region)
	     ;; C-u C-c C-c 
	     (define-key go-mode-map "\C-u\C-c\C-c" 'uncomment-region)
             ))

(setq gofmt-command "goimports")
(add-hook 'before-save-hook #'gofmt-before-save)

(require 'go-eldoc)
(add-hook 'go-mode-hook 'go-eldoc-setup)
(set-face-attribute 'eldoc-highlight-function-argument nil
                    :underline t :foreground "darkgreen"
                    :weight 'bold)
