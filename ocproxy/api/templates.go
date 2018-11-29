package api

var publicLinkTemplate string = `
<!DOCTYPE html>
<html class="ng-csp" data-placeholder-focus="false" lang="en" >
	<head>
		<meta charset="utf-8">
		<title> CERNBox</title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="referrer" content="never">
		<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0">
		<meta name="theme-color" content="#1d2d44">

		<link rel="icon" href="/{{ .BaseUrl }}/core/img/favicon.ico">
		<link rel="apple-touch-icon-precomposed" href="/{{ .BaseUrl }}/core/img/favicon-touch.png">
		<link rel="mask-icon" sizes="any" href="/{{ .BaseUrl }}/core/img/favicon-mask.svg" color="#1d2d44">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/styles.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/inputs.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/header.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/icons.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fonts.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/apps.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/global.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fixes.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/multiselect.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/mobile.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/select2/select2.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/jquery-ui/themes/base/jquery-ui.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery-ui-fixes.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/tooltip.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/share.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_versions/css/versions.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery.ocdialog.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_sharing/css/public.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_sharing/css/mobile.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files/css/files.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files/css/upload.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/cernbox-theme/core/css/styles.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/wopiviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/gallery/css/slideshow.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/gallery/css/gallerybutton.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_pdfviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/vendor/JSRootPainter.min.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/vendor/JSRootInterface.min.css">
		<script src="/{{ .BaseUrl }}/index.php/core/js/oc.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery/dist/jquery.min.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery-migrate/jquery-migrate.min.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery-ui/ui/jquery-ui.custom.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/underscore/underscore.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/moment/min/moment-with-locales.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/handlebars/handlebars.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/blueimp-md5/js/md5.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/bootstrap/js/tooltip.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/backbone/backbone.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/es6-promise/dist/es6-promise.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/davclient.js/lib/client.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/clipboard/dist/clipboard.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/bowser/src/bowser.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/jquery.ocdialog.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-dialogs.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/js.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/l10n.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/octemplate.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/eventsource.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/config.js"></script>
		<script src="/{{ .BaseUrl }}/core/search/js/search.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-requesttoken.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/apps.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/mimetype.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/mimetypelist.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/snapjs/dist/latest/snap.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-backbone.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/backgroundjobs.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/shareconfigmodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharemodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharescollection.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/shareitemmodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogresharerinfoview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinklistview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinkshareview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogmailview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinksocialview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogexpirationview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogshareelistview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/share.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/files/fileinfo.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/files/client.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/file-upload.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files_sharing/js/public.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileactions.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileactionsmenu.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/jquery.fileupload.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/filesummary.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/breadcrumb.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileinfomodel.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/newfilemenu.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/files.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/filelist.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/keyboardshortcuts.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/vendor/bigshot/bigshot-compressed.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/galleryutility.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/galleryfileaction.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/slideshow.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/slideshowcontrols.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/slideshowzoomablepreview.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/gallerybutton.js"></script>
		<script src="/{{ .BaseUrl }}/apps/gallery/js/vendor/dompurify/src/purify.js"></script>

		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/script.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootCore.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/d3.v3.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootPainter.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootInterface.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootPainter.jquery.min.js"></script>
		<script>
		(function ($, OC) {

			$(document).ready(function () {
				var data = $("data[key='cernboxauthtoken']");
				var accessToken = data.attr('x-access-token');
				if(accessToken) {
					OC["X-Access-Token"] = accessToken;
					/*
					OC.Files.getClient()["_defaultHeaders"]["X-Access-Token"] = accessToken;

					$.ajaxSetup({
						    headers: { 'X-Access-Token': accessToken }
					});

					$(document).on('ajaxSend',function(elm, xhr, settings) {
						xhr.setRequestHeader('X-Access-Token', accessToken);
					});
					*/

					XMLHttpRequest.prototype.origOpen = XMLHttpRequest.prototype.open;
					XMLHttpRequest.prototype.open   = function () {
						this.origOpen.apply(this, arguments);
						this.setRequestHeader('X-Access-Token', accessToken);
					};

				}
			});

		})(jQuery, OC);
		</script>
    		<script src="/{{ .BaseUrl }}/apps/wopiviewer/js/script.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files_pdfviewer/js/previewplugin.js"></script>
	<body id="body-public">
	<data key="cernboxauthtoken" x-access-token="{{ .AccessToken }}" />
	<noscript>
	<div id="nojavascript">
	<div> This application requires JavaScript for correct operation. Please <a href="http://enable-javascript.com/" target="_blank" rel="noreferrer">enable JavaScript</a> and reload the page.</div>
	</div>
	</noscript>

	<div id="notification-container">
		<div id="notification" style="display: none;"></div>
	</div>

	<input type="hidden" id="filesApp" name="filesApp" value="1">
	<input type="hidden" id="isPublic" name="isPublic" value="1">
	<input type="hidden" name="dir" value="/" id="dir">
	<input type="hidden" name="downloadURL" value="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}" id="downloadURL">
	<input type="hidden" name="sharingToken" value="{{ .Token }}" id="sharingToken">
	<input type="hidden" name="filename" value="/" id="filename">
	<input type="hidden" name="mimetype" value="httpd/unix-directory" id="mimetype">
	<input type="hidden" name="previewSupported" value="false" id="previewSupported">
	<input type="hidden" name="mimetypeIcon" value="/core/img/filetypes/folder.svg" id="mimetypeIcon">
	<input type="hidden" name="filesize" value="28" id="filesize">
	<input type="hidden" name="maxSizeAnimateGif" value="10" id="maxSizeAnimateGif">

	<header>
		<div id="header" class="share-folder" data-protected="false"
			 data-owner-display-name="admin" data-owner="admin" data-name="Test folder">
			<a href="/{{ .BaseUrl }}/index.php" title="" id="owncloud">
				<h1 class="logo-icon"></h1>
			</a>

			<div id="logo-claim" style="display:none;"></div>
					<div class="header-right">
				<span id="details">
					<a href="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}" id="download" class="button">
						<img class="svg" alt="" src="/{{ .BaseUrl }}/core/img/actions/download.svg"/>
						<span id="download-text">Download</span>
					</a>
				</span>
			</div>
				</div>
	</header>
	<div id="content-wrapper">
		<div id="content">
			<div id="preview">
								<div id="controls">
			<div class="actions creatable hidden">
				<div id="uploadprogresswrapper">
					<div id="uploadprogressbar">
						<em class="label outer" style="display:none"><span class="desktop">Uploading...</span><span class="mobile">...</span></em>
					</div>
					<input type="button" class="stop icon-close" style="display:none" value="" />
				</div>
			</div>
			<div id="file_action_panel"></div>
			<div class="notCreatable notPublic hidden">
				You don't have permission to upload or create files here</div>
			<input type="hidden" name="permissions" value="" id="permissions">
		<input type="hidden" id="free_space" value="INF">
			<input type="hidden" id="publicUploadRequestToken" name="requesttoken" value="PTUadAJrOx8NOxZGNCEpEQtoAA0wQBMbCWcRT3ROFB8=:WcL6HZXeYsbiDDAfGGyhq9jJe7V7B7Xi3kzHJUmTs/g=" />
		<input type="hidden" id="dirToken" name="dirToken" value="{{ .Token }}" />
			<input type="hidden" class="max_human_file_size"
			   value="(max INF PB)">
	</div>

	<div id="emptycontent" class="hidden">
		<div class="icon-folder"></div>
		<h2>No files in here</h2>
		<p class="uploadmessage hidden">Upload some content or sync with your devices!</p>
	</div>

	<div class="nofilterresults emptycontent hidden">
		<div class="icon-search"></div>
		<h2>No entries found in this folder</h2>
		<p></p>
	</div>

	<table id="filestable" data-allow-public-upload="no" data-preview-x="32" data-preview-y="32">
		<thead>
			<tr>
				<th id='headerName' class="hidden column-name">
					<div id="headerName-container">
						<input type="checkbox" id="select_all_files" class="select-all checkbox"/>
						<label for="select_all_files">
							<span class="hidden-visually">Select all</span>
						</label>
						<a class="name sort columntitle" data-sort="name"><span>Name</span><span class="sort-indicator"></span></a>
						<span id="selectedActionsList" class="selectedActions">
							<a href="" class="download">
								<span class="icon icon-download"></span>
								<span>Download</span>
							</a>
						</span>
					</div>
				</th>
				<th id="headerSize" class="hidden column-size">
					<a class="size sort columntitle" data-sort="size"><span>Size</span><span class="sort-indicator"></span></a>
				</th>
				<th id="headerDate" class="hidden column-mtime">
					<a id="modified" class="columntitle" data-sort="mtime"><span>Modified</span><span class="sort-indicator"></span></a>
						<span class="selectedActions"><a href="" class="delete-selected">
							<span>Delete</span>
							<span class="icon icon-delete"></span>
						</a></span>
				</th>
			</tr>
		</thead>
		<tbody id="fileList">
		</tbody>
		<tfoot>
		</tfoot>
	</table>
	<input type="hidden" name="dir" id="dir" value="" />
	<div class="hiddenuploadfield">
		<input type="file" id="file_upload_start" class="hiddenuploadfield" name="files[]" />
	</div>
	<div id="editor"></div><!-- FIXME Do not use this div in your app! It is deprecated and will be removed in the future! -->
	<div id="uploadsize-message" title="Upload too large">
		<p>
		The files you are trying to upload exceed the maximum size for file uploads on this server.	</p>
	</div>
						</div>
		</div>
		<footer>
			<p class="info">
				<a href="https://cernbox.web.cern.ch" target="_blank" rel="noreferrer">CERNBox</a> &ndash; {{ .Note }}</p>
		</footer>
	</div>
	</body>
</html>
`

var publicLinkTemplatePassword = `
<!DOCTYPE html>
<html class="ng-csp" data-placeholder-focus="false" lang="en" >
  <head data-requesttoken="GxEhFD47HA8bZlM1MyQfGChWczklBDM4KyQVO18ATHE=:hvNGFqhLL7cYAVIiqo+qr1KMyTWijp62F95Hkn7Nxs0=">
    <meta charset="utf-8">
    <title>
      CERNBox		
    </title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="referrer" content="never">
    <meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0">
    <meta name="apple-itunes-app" content="app-id=543672169">
    <meta name="theme-color" content="#1d2d44">
    <link rel="icon" href="/{{ .BaseUrl }}/core/img/favicon.ico">
    <link rel="apple-touch-icon-precomposed" href="/{{ .BaseUrl }}/core/img/favicon-touch.png">
    <link rel="mask-icon" sizes="any" href="/{{ .BaseUrl }}/core/img/favicon-mask.svg" color="#1d2d44">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/select2/select2.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/styles.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/inputs.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/header.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/icons.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fonts.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/apps.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/global.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fixes.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/multiselect.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/mobile.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/jquery-ui/themes/base/jquery-ui.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery-ui-fixes.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/tooltip.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/share.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/apps/cernbox-theme/core/css/styles.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_versions/css/versions.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery.ocdialog.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_sharing/css/authenticate.css">
    <script src="/{{ .BaseUrl }}/index.php/core/js/oc.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery/dist/jquery.min.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery-migrate/jquery-migrate.min.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery-ui/ui/jquery-ui.custom.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/underscore/underscore.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/moment/min/moment-with-locales.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/handlebars/handlebars.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/blueimp-md5/js/md5.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/bootstrap/js/tooltip.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/backbone/backbone.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/es6-promise/dist/es6-promise.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/davclient.js/lib/client.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/clipboard/dist/clipboard.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/bowser/src/bowser.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/jquery.ocdialog.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-dialogs.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/js.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/l10n.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/octemplate.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/eventsource.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/config.js"></script>
    <script src="/{{ .BaseUrl }}/core/search/js/search.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-requesttoken.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/apps.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/mimetype.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/mimetypelist.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/snapjs/dist/latest/snap.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/select2/select2.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-backbone.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/placeholder.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/jquery.avatar.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/backgroundjobs.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/shareconfigmodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharemodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharescollection.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/shareitemmodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogresharerinfoview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinklistview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinkshareview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogmailview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinksocialview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogexpirationview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogshareelistview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/share.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/files/fileinfo.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/files/client.js"></script>
    <script src="/{{ .BaseUrl }}/apps/files_sharing/js/authenticate.js"></script>
    <script src="/{{ .BaseUrl }}/apps/wopiviewer/js/script.js"></script>
  </head>
  <body id="body-login">
    <noscript>
      <div id="nojavascript">
        <div>
          This application requires JavaScript for correct operation. Please <a href="http://enable-javascript.com/" target="_blank" rel="noreferrer">enable JavaScript</a> and reload the page.		
        </div>
      </div>
    </noscript>
    <div class="wrapper">
      <div class="v-align">
        <header role="banner">
          <div id="header">
            <div class="logo">
              <h1 class="hidden-visually">
              </h1>
            </div>
            <div id="logo-claim" style="display:none;"></div>
          </div>
        </header>
        <form method="post">
          <fieldset>
            <div class="warning-info">This share is password-protected</div>
            <p>
              <label for="password" class="infield">Password</label>
              <input type="hidden" name="requesttoken" value="GxEhFD47HA8bZlM1MyQfGChWczklBDM4KyQVO18ATHE=:hvNGFqhLL7cYAVIiqo+qr1KMyTWijp62F95Hkn7Nxs0=" />
              <input type="password" name="password" id="password"
                placeholder="Password" value=""
                autocomplete="off" autocapitalize="off" autocorrect="off"
                autofocus />
              <input type="submit" id="password-submit" 
                class="svg icon-confirm input-button-inline" value="" disabled="disabled" />
            </p>
          </fieldset>
        </form>
        <div class="push"></div>
        <!-- for sticky footer -->
      </div>
    </div>
    <footer role="contentinfo">
      <p class="info">
        <a href="https://cernbox.web.cern.ch" target="_blank" rel="noreferrer">CERNBox</a> &ndash; The CERN Cloud Storage</p>
    </footer>
  </body>
</html>
`

var publicLinkTemplateNotFound string = `
<!DOCTYPE html>
<html class="ng-csp" data-placeholder-focus="false" lang="en" >
  <head data-requesttoken="bwE2RRwgGDM2KRF9OwpBURgTVVlxJlojGy9kDghjIw4=:VMwtwFJuWYP7hl57ijm/CB3zOn0o01fZRmuxPp/xxTI=">
    <meta charset="utf-8">
    <title>
      CERNBox		
    </title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="referrer" content="never">
    <meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0">
    <meta name="apple-itunes-app" content="app-id=543672169">
    <meta name="theme-color" content="#1d2d44">
    <link rel="icon" href="/{{ .BaseUrl }}/core/img/favicon.ico">
    <link rel="apple-touch-icon-precomposed" href="/{{ .BaseUrl }}/core/img/favicon-touch.png">
    <link rel="mask-icon" sizes="any" href="/{{ .BaseUrl }}/core/img/favicon-mask.svg" color="#1d2d44">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/select2/select2.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/styles.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/inputs.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/header.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/icons.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fonts.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/apps.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/global.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fixes.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/multiselect.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/mobile.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/jquery-ui/themes/base/jquery-ui.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery-ui-fixes.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/tooltip.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/share.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_versions/css/versions.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery.ocdialog.css">
    <link rel="stylesheet" href="/{{ .BaseUrl }}/apps/cernbox-theme/core/css/styles.css">
    <script src="/{{ .BaseUrl }}/index.php/core/js/oc.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery/dist/jquery.min.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery-migrate/jquery-migrate.min.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/jquery-ui/ui/jquery-ui.custom.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/underscore/underscore.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/moment/min/moment-with-locales.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/handlebars/handlebars.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/blueimp-md5/js/md5.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/bootstrap/js/tooltip.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/backbone/backbone.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/es6-promise/dist/es6-promise.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/davclient.js/lib/client.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/clipboard/dist/clipboard.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/bowser/src/bowser.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/jquery.ocdialog.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-dialogs.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/js.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/l10n.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/octemplate.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/eventsource.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/config.js"></script>
    <script src="/{{ .BaseUrl }}/core/search/js/search.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-requesttoken.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/apps.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/mimetype.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/mimetypelist.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/snapjs/dist/latest/snap.js"></script>
    <script src="/{{ .BaseUrl }}/core/vendor/select2/select2.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/oc-backbone.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/placeholder.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/jquery.avatar.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/backgroundjobs.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/shareconfigmodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharemodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharescollection.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/shareitemmodel.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogresharerinfoview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinklistview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinkshareview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogmailview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialoglinksocialview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogexpirationview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogshareelistview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/sharedialogview.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/share.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/files/fileinfo.js"></script>
    <script src="/{{ .BaseUrl }}/core/js/files/client.js"></script>
  </head>
  <body id="body-login">
    <noscript>
      <div id="nojavascript">
        <div>
          This application requires JavaScript for correct operation. Please <a href="http://enable-javascript.com/" target="_blank" rel="noreferrer">enable JavaScript</a> and reload the page.		
        </div>
      </div>
    </noscript>
    <div class="wrapper">
      <div class="v-align">
        <header role="banner">
          <div id="header">
            <div class="logo">
              <h1 class="hidden-visually">
              </h1>
            </div>
            <div id="logo-claim" style="display:none;"></div>
          </div>
        </header>
        <ul>
          <li class="error">
            File not found or expired<br>
            <p class="hint">The specified document has not been found on the server or it has expired.</p>
            <p class="hint"><a href="/{{ .BaseUrl }}/index.php">You can click here to return to CERNBox.</a></p>
          </li>
        </ul>
        <div class="push"></div>
        <!-- for sticky footer -->
      </div>
    </div>
    <footer role="contentinfo">
      <p class="info">
        <a href="https://cernbox.web.cern.ch" target="_blank" rel="noreferrer">CERNBox</a> The CERN Cloud Storage</p>
    </footer>
  </body>
</html>
`

var publicLinkTemplateFile = `
<!DOCTYPE html>
<html class="ng-csp" data-placeholder-focus="false" lang="en" >
	<head>
		<meta charset="utf-8">
		<title> CERNBox</title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="referrer" content="never">
		<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0">
		<meta name="theme-color" content="#1d2d44">

		<link rel="icon" href="/{{ .BaseUrl }}/core/img/favicon.ico">
		<link rel="apple-touch-icon-precomposed" href="/{{ .BaseUrl }}/core/img/favicon-touch.png">
		<link rel="mask-icon" sizes="any" href="/{{ .BaseUrl }}/core/img/favicon-mask.svg" color="#1d2d44">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/styles.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/inputs.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/header.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/icons.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fonts.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/apps.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/global.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/fixes.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/multiselect.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/mobile.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/select2/select2.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/vendor/jquery-ui/themes/base/jquery-ui.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery-ui-fixes.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/tooltip.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/share.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_versions/css/versions.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/core/css/jquery.ocdialog.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_sharing/css/public.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_sharing/css/mobile.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files/css/files.css">
		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files/css/upload.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/cernbox-theme/core/css/styles.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/wopiviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/files_pdfviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/style.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/vendor/JSRootPainter.min.css">
    		<link rel="stylesheet" href="/{{ .BaseUrl }}/apps/rootviewer/css/vendor/JSRootInterface.min.css">
		<script src="/{{ .BaseUrl }}/index.php/core/js/oc.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery/dist/jquery.min.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery-migrate/jquery-migrate.min.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/jquery-ui/ui/jquery-ui.custom.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/underscore/underscore.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/moment/min/moment-with-locales.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/handlebars/handlebars.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/blueimp-md5/js/md5.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/bootstrap/js/tooltip.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/backbone/backbone.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/es6-promise/dist/es6-promise.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/davclient.js/lib/client.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/clipboard/dist/clipboard.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/bowser/src/bowser.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/jquery.ocdialog.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-dialogs.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/js.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/l10n.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/octemplate.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/eventsource.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/config.js"></script>
		<script src="/{{ .BaseUrl }}/core/search/js/search.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-requesttoken.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/apps.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/mimetype.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/mimetypelist.js"></script>
		<script src="/{{ .BaseUrl }}/core/vendor/snapjs/dist/latest/snap.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/oc-backbone.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/backgroundjobs.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/shareconfigmodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharemodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharescollection.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/shareitemmodel.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogresharerinfoview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinklistview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinkshareview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogmailview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialoglinksocialview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogexpirationview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogshareelistview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/sharedialogview.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/share.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/files/fileinfo.js"></script>
		<script src="/{{ .BaseUrl }}/core/js/files/client.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/file-upload.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files_sharing/js/public.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileactions.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileactionsmenu.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/jquery.fileupload.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/filesummary.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/breadcrumb.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/fileinfomodel.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/newfilemenu.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/files.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/filelist.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files/js/keyboardshortcuts.js"></script>

		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/script.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootCore.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/d3.v3.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootPainter.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootInterface.min.js"></script>
		<script src="/{{ .BaseUrl }}/apps/rootviewer/js/scripts/JSRootPainter.jquery.min.js"></script>
		<script>
		(function ($, OC) {
			$(document).ready(function () {
				var data = $("data[key='cernboxauthtoken']");
				var accessToken = data.attr('x-access-token');
				if(accessToken) {
					OC["X-Access-Token"] = accessToken;
					/*
					OC.Files.getClient()["_defaultHeaders"]["X-Access-Token"] = accessToken;

					$.ajaxSetup({
						    headers: { 'X-Access-Token': accessToken }
					});

					$(document).on('ajaxSend',function(elm, xhr, settings) {
						xhr.setRequestHeader('X-Access-Token', accessToken);
					});
					*/

					XMLHttpRequest.prototype.origOpen = XMLHttpRequest.prototype.open;
					XMLHttpRequest.prototype.open   = function () {
						this.origOpen.apply(this, arguments);
						this.setRequestHeader('X-Access-Token', accessToken);
					};

				}
			});
		})(jQuery, OC);
		</script>
    		<script src="/{{ .BaseUrl }}/apps/wopiviewer/js/script.js"></script>
		<script src="/{{ .BaseUrl }}/apps/files_pdfviewer/js/previewplugin.js"></script>
		{{ if (eq .Mime "application/pynb") }}
    		<script src="/{{ .BaseUrl }}/apps/swanviewer/js/script.js"></script>
		{{ end }}
	<body id="body-public">
	<data key="cernboxauthtoken" x-access-token="{{ .AccessToken }}" />
	<noscript>
	<div id="nojavascript">
	<div> This application requires JavaScript for correct operation. Please <a href="http://enable-javascript.com/" target="_blank" rel="noreferrer">enable JavaScript</a> and reload the page.</div>
	</div>
	</noscript>

	<div id="notification-container">
		<div id="notification" style="display: none;"></div>
	</div>

	<input type="hidden" id="filesApp" name="filesApp" value="1">
	<input type="hidden" id="isPublic" name="isPublic" value="1">
	<input type="hidden" name="dir" value="/" id="dir">
	<input type="hidden" name="downloadURL" value="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}" id="downloadURL">
	<input type="hidden" name="sharingToken" value="{{ .Token }}" id="sharingToken">
	<input type="hidden" name="filename" value="/" id="filename">
	<input type="hidden" name="mimetype" value="{{ .Mime }}" id="mimetype">
	<input type="hidden" name="previewSupported" value="false" id="previewSupported">
	<input type="hidden" name="mimetypeIcon" value="/core/img/filetypes/folder.svg" id="mimetypeIcon">
	<input type="hidden" name="filesize" value="28" id="filesize">
	<input type="hidden" name="maxSizeAnimateGif" value="10" id="maxSizeAnimateGif">
	<script>
			// change mimetype
			$("#mimetypeIcon").val(OC.MimeType.getIconUrl($("#mimetype").val()));
	</script>
	<header>
		<div id="header" class="share-folder" data-protected="false"
			 data-owner-display-name="admin" data-owner="admin" data-name="Test folder">
			<a href="/{{ .BaseUrl }}/index.php" title="" id="owncloud">
				<h1 class="logo-icon"></h1>
			</a>

			<div id="logo-claim" style="display:none;"></div>
					<div class="header-right">
				<span id="details">
					<a href="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}" id="download" class="button">
						<img class="svg" alt="" src="/{{ .BaseUrl }}/core/img/actions/download.svg"/>
						<span id="download-text">Download</span>
					</a>
				</span>
			</div>
				</div>
	</header>
    <div id="content-wrapper">
      <div id="content">
        <div id="preview">
          <!-- Preview frame is filled via JS to support SVG images for modern browsers -->
          <div id="imgframe"></div>
          <div class="directDownload">
            <a href="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}" id="downloadFile" class="button">
            <img class="svg" alt="" src="/{{ .BaseUrl }}/core/img/actions/download.svg"/>
            Download {{ .ShareName }}
            <!--Download {{ .ShareName }} (4.7 MB)-->
            </a>
          </div>
          <div class="directLink">
            <label for="directLink">Direct link</label>
            <input id="directLink" type="text" readonly value="https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download?x-access-token={{ .AccessToken }}">
          </div>

	  <!--
	  {{ if (eq .Mime "application/pynb") }}
	  <a href="https://cern.ch/swanserver/cgi-bin/go?projurl=https://{{ .OverwriteHost }}/index.php/s/{{ .Token }}/download%3Fx-access-token={{ .AccessToken }}" target="_blank"><img class="svg" alt="" src="/{{ .BaseUrl }}/apps/swanviewer/img//badge_swan_white_150.svg"></a>
	  {{ end }}
	  -->
        </div>
      </div>
      <footer>
        <p class="info">
          <a href="https://cernbox.wen.cern.ch" target="_blank" rel="noreferrer">CERNBox</a> &ndash; The CERN Cloud Storage		
        </p>
      </footer>
    </div>
	</body>
</html>
`

var notebookScript = `
#!/usr/bin/env python3

import nbformat
from nbconvert import HTMLExporter
import sys
 
source_fn = sys.argv[1]
target_fn = sys.argv[2]


with open(source_fn, 'r') as content_file:
        content = content_file.read()
        tnb = nbformat.reads(content, as_version=4)
        html_exporter = HTMLExporter()
        html_exporter.template_file = 'full' # basic html to be embeded in CERNBox
        (body, resources) = html_exporter.from_notebook_node(tnb)
        with open(target_fn, "w") as fd:
                fd.write(body)
# Installation: pip3 install nbconvert
`
