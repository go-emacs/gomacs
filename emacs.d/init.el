(set-language-environment "Japanese")

;; utf-8
(prefer-coding-system 'utf-8)
(setq coding-system-for-read 'utf-8)
(setq coding-system-for-write 'utf-8)

(menu-bar-mode 0)
(if (display-graphic-p)
    (progn
      (tool-bar-mode 0)))
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
(global-set-key "\C-_" 'advertised-undo)
(define-key minibuffer-local-completion-map "\C-w" 'backward-kill-word)

(add-to-list 'load-path (concat gomacs-emacsd-path "/elisp"))
(when (<= emacs-major-version 24.2)
  (require 'cl-lib))

(require 'go-mode-autoloads)
(require 'golint)

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

(load (concat gopath "/src/golang.org/x/tools/cmd/oracle/oracle.el"))
(setq go-oracle-command (concat gopath "/bin/oracle"))
(add-hook 'go-mode-hook (lambda () (go-oracle-mode t) ))

(load (concat gopath "/src/golang.org/x/tools/refactor/rename/rename.el"))

(load (concat gopath "/src/github.com/nsf/gocode/emacs/go-autocomplete.el"))
(add-to-list 'ac-dictionary-directories (concat gopath "/src/github.com/auto-complete/auto-complete/dict"))
(require 'go-autocomplete)
(require 'auto-complete-config)
(ac-config-default)
(ac-set-trigger-key "TAB")

(require 'yasnippet)
(add-to-list 'yas-snippet-dirs (concat gopath "/src/github.com/atotto/yasnippet-golang"))
(yas-global-mode 1)

(require 'autoinsert)
(setq auto-insert-directory (concat gomacs-emacsd-path "/_templates/"))
;(define-auto-insert "\\.go\\'" "T.go")
;(define-auto-insert "\\test.go\\'" "T_test.go")
(add-hook 'find-file-hooks 'auto-insert)

(setq compilation-window-height 10)
