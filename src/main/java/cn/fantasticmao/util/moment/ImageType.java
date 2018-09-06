package cn.fantasticmao.util.moment;

import java.util.stream.Stream;

/**
 * ImageType
 *
 * @author maodh
 * @since 2018/9/5
 */
enum ImageType {
    BMP(".bmp", "424D"),
    JPG(".jpg", "FFD8FFDB"),
    JPEG(".jpeg", "FFD8FF"),
    PNG(".png", "89504E47");

    public final String extension;
    public final String magicNumber;

    ImageType(String extension, String magicNumber) {
        this.extension = extension;
        this.magicNumber = magicNumber;
    }

    public static ImageType ofExtension(final String extension) {
        return Stream.of(ImageType.values())
                .filter(fileType -> extension.equals(fileType.extension))
                .findFirst().orElse(null);
    }
}
